package services

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
)

type frService struct {
	frRepo  repository.FriendRequestRepository
	fRepo   repository.FriendRepository
	frlRepo repository.FriendRequestLogRepository
	cSvc    service.ConversationService
}

func (s *frService) SendFriendRequest(
	ctx context.Context,
	userID string,
	dto dto.SendFriendRequestDto,
) error {
	fr := &models.FriendRequest{
		SenderId:   userID,
		ReceiverId: dto.ReceiverId,
	}

	//check already send friend request
	pending := s.frRepo.HasPendingRequest(ctx, fr.SenderId, fr.ReceiverId)

	if pending {
		return &e.ConflictError{Message: "Hang tight! Friend request already sent."}
	}

	//check already friend
	exist := s.fRepo.AlreadyFriends(ctx, fr.SenderId, fr.ReceiverId)

	if exist {
		return &e.ConflictError{Message: "You're already connected."}
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
		canSend, err := s.frlRepo.HasRejectedFriendRequestLog(frl)

		if err != nil {
			return err
		}
		if !canSend {
			return &e.BadRequestError{Message: "You can't send this right now!"}
		}
		err = s.frlRepo.CreateFriendRequestLog(frl)

		if err != nil {
			return &e.InternalServerError{Message: "Something went wrong. Please try agian laster!"}
		}
		//make friend request
		return s.frRepo.SendFriendRequest(fr)
	}

	return nil

}

func (s *frService) DecideFriendRequest(
	ctx context.Context,
	userID string,
	dfrDto dto.DecideFriendRequestDto,
) error {

	dfr := &models.FriendRequest{
		ID:         dfrDto.FriendRequestId,
		ReceiverId: userID,
		Status:     dfrDto.Status,
	}
	//check decide status is ACCEPTED
	//then delete the friend request row
	//write into friends database for both friendship 2 user id

	fr, err := s.frRepo.GetFriendRequestByID(ctx, dfr.ID)

	if err != nil {
		return &e.NotFoundError{Message: "Friend Request Not Found!"}
	}

	//current user is equal to receiver
	if fr.ReceiverId == dfr.ReceiverId {
		if models.FriendRequestAccepted == dfr.Status {

			err := s.frRepo.DeleteFriendRequestByID(fr.ID)

			if err != nil {
				return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
			}

			//insert two row [bidirectional]
			friends := make([]*models.Friend, 2)
			friends[0] = &models.Friend{
				UserID:   fr.SenderId,
				FriendID: fr.ReceiverId,
			}

			friends[1] = &models.Friend{
				UserID:   fr.ReceiverId,
				FriendID: fr.SenderId,
			}
			//make friendship [bidirectional]
			for _, f := range friends {
				if err := s.fRepo.CreateFriendShip(f); err != nil {
					return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
				}
			}

			//make Action to ACCEPTED
			acceptFrl := &models.FriendRequestLog{
				SenderID:    fr.SenderId,
				ReceiverID:  dfr.ReceiverId,
				Action:      models.ActionAccepted,
				PerformedBy: dfr.ReceiverId,
			}

			canSend, err := s.frlRepo.HasRejectedFriendRequestLog(acceptFrl)

			if err != nil {
				return err
			}

			if !canSend {
				return &e.BadRequestError{Message: "You can't send this right now!"}
			}

			err = s.frlRepo.CreateFriendRequestLog(acceptFrl)

			if err != nil {
				return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
			}

			//Make for conversation and conversation member for both after ACCEPTED
			cDto := dto.CreateConversationDto{
				IsGroup:  false,
				Name:     nil,
				MemberID: &[]string{fr.SenderId, fr.ReceiverId},
			}
			if err := s.cSvc.CreateConversation(nil, cDto); err != nil {
				return err
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

			err = s.frlRepo.CreateFriendRequestLog(rejectFrl)

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

func (s *frService) GetAllFriendRequestLog(ctx context.Context, userID string) ([]response.FriendRequestResponse, error) {
	logs, err := s.frRepo.GetAllFriendRequestLog(ctx, userID)

	if err != nil {
		return nil, &e.InternalServerError{Message: err.Error()}
	}
	return logs, nil
}
