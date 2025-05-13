package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/pkg/common"
)

type FriendRequestService struct {
	frRepo  repository.FriendRequestRepository
	fRepo   repository.FriendRepository
	frlRepo repository.FriendRequestLogRepository
}

func NewFriendRequestService(
	frRepo repository.FriendRequestRepository,
	fRepo repository.FriendRepository,
	frlRepo repository.FriendRequestLogRepository,
) *FriendRequestService {
	return &FriendRequestService{
		frRepo:  frRepo,
		fRepo:   fRepo,
		frlRepo: frlRepo,
	}
}

func (s *FriendRequestService) SendFriendRequest(fr *models.FriendRequest) error {

	//check already send friend request
	pending := s.frRepo.HasPendingRequest(fr.SenderId, fr.ReceiverId)

	if pending {
		return &common.ConflictError{Message: "Friend request already send!"}
	}

	//check already friend
	exist := s.frRepo.AlreadyFriends(fr.SenderId, fr.ReceiverId)

	if exist {
		return &common.ConflictError{Message: "Already friend each other!"}
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
			return &common.InternalServerError{Message: "Something went wrong. Please try agian laster!"}
		}
		//make friend request
		return s.frRepo.SendFriendRequest(fr)
	}

	return nil

}

func (s *FriendRequestService) DecideFriendRequest(dfr *models.FriendRequest) error {

	//check decide status is ACCEPTED
	//then delete the friend request row
	//write into friends database for both friendship 2 user id

	fr, err := s.frRepo.FindById(dfr.ID)

	if err != nil {
		return &common.NotFoundError{Message: "Friend Request Not Found!"}
	}

	//current user is equal to receiver
	if fr.ReceiverId == dfr.ReceiverId {
		if models.StatusAccepted == dfr.Status {

			err := s.frRepo.DeleteById(fr.ID)

			if err != nil {
				return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
			}

			f := &models.Friend{
				UserID:   fr.SenderId,
				FriendID: fr.ReceiverId,
			}

			//make friendship
			s.fRepo.CreateFriendShip(f)

			//make Action to ACCEPTED
			acceptFrl := &models.FriendRequestLog{
				SenderID:    dfr.ReceiverId,
				ReceiverID:  dfr.ReceiverId,
				Action:      models.ActionAccepted,
				PerformedBy: dfr.ReceiverId,
			}

			err = s.frlRepo.CreateFriendRequestLog(acceptFrl)

			if err != nil {
				return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
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
				return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
			}

			return s.frRepo.RejectFriendRequest(dfr)
		}
	} else {
		return &common.ForbiddenError{Message: "You are not allowed to do this action!"}
	}
	return nil
}
