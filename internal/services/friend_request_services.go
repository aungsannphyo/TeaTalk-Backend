package services

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/gin-gonic/gin"
)

type frService struct {
	frRepo  repository.FriendRequestRepository
	fRepo   repository.FriendRepository
	frlRepo repository.FriendRequestLogRepository
}

func (s *frService) SendFriendRequest(c *gin.Context, dto dto.SendFriendRequestDto) error {
	fr := &models.FriendRequest{
		SenderId:   c.GetString("userId"),
		ReceiverId: dto.ReceiverId,
	}

	//check already send friend request
	pending := s.frRepo.HasPendingRequest(c.Request.Context(), fr.SenderId, fr.ReceiverId)

	if pending {
		return &e.ConflictError{Message: "Friend request already send!"}
	}

	//check already friend
	exist := s.fRepo.AlreadyFriends(c.Request.Context(), fr.SenderId, fr.ReceiverId)

	if exist {
		return &e.ConflictError{Message: "Already friend each other!"}
	}

	if !pending && !exist {

		//make friend request log by default action SENT
		frl := &models.FriendRequestLog{
			SenderID:    fr.SenderId,
			ReceiverID:  fr.ReceiverId,
			Action:      models.ActionSend,
			PerformedBy: fr.SenderId,
		}
		//default action SENT
		err := s.frlRepo.CreateFriendRequestLog(frl)
		if err != nil {
			return &e.InternalServerError{Message: "Something went wrong. Please try agian laster!"}
		}
		//make friend request
		return s.frRepo.SendFriendRequest(fr)
	}

	return nil

}

func (s *frService) DecideFriendRequest(c *gin.Context, dto dto.DecideFriendRequestDto) error {

	dfr := &models.FriendRequest{
		ID:         dto.FriendRequestId,
		ReceiverId: c.GetString("userId"),
		Status:     dto.Status,
	}

	//check decide status is ACCEPTED
	//then delete the friend request row
	//write into friends database for both friendship 2 user id

	fr, err := s.frRepo.FindById(c.Request.Context(), dfr.ID)

	if err != nil {
		return &e.NotFoundError{Message: "Friend Request Not Found!"}
	}

	//current user is equal to receiver
	if fr.ReceiverId == dfr.ReceiverId {
		if models.FriendRequestAccepted == dfr.Status {

			err := s.frRepo.DeleteById(fr.ID)

			if err != nil {
				return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
			}

			//insert two row [bidirectional]
			f1 := &models.Friend{
				UserID:   fr.SenderId,
				FriendID: fr.ReceiverId,
			}

			f2 := &models.Friend{
				UserID:   fr.ReceiverId,
				FriendID: fr.SenderId,
			}
			//make friendship [bidirectional]
			errf1 := s.fRepo.CreateFriendShip(f1)
			errf2 := s.fRepo.CreateFriendShip(f2)

			if errf1 != nil || errf2 != nil {
				return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
			}

			//make Action to ACCEPTED
			acceptFrl := &models.FriendRequestLog{
				SenderID:    dfr.ReceiverId,
				ReceiverID:  dfr.ReceiverId,
				Action:      models.ActionAccepted,
				PerformedBy: dfr.ReceiverId,
			}

			err = s.frlRepo.CreateFriendRequestLog(acceptFrl)

			if err != nil {
				return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
			}
		} else {
			//Reject Case
			//make Action to REJECTED
			rejectFrl := &models.FriendRequestLog{
				SenderID:    dfr.ReceiverId,
				ReceiverID:  dfr.ReceiverId,
				Action:      models.ActionRejected,
				PerformedBy: dfr.ReceiverId,
			}

			err := s.frlRepo.CreateFriendRequestLog(rejectFrl)

			if err != nil {
				return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
			}

			return s.frRepo.RejectFriendRequest(dfr)
		}
	} else {
		return &e.ForbiddenError{Message: "You are not allowed to do this action!"}
	}
	return nil
}
