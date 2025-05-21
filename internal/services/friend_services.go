package services

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
)

type fService struct {
	fRepo   repository.FriendRepository
	frRepo  repository.FriendRequestRepository
	frlRepo repository.FriendRequestLogRepository
}

func (s *fService) MakeUnFriend(userID string, dto dto.UnFriendDto) error {

	user := &models.Friend{
		UserID:   userID,
		FriendID: dto.FriendID,
	}

	friend := &models.Friend{
		UserID:   dto.FriendID,
		FriendID: userID,
	}

	frl := &models.FriendRequestLog{
		SenderID:    userID,
		ReceiverID:  dto.FriendID,
		Action:      models.ActionUnFriended,
		PerformedBy: userID,
	}

	canSend, err := s.frlRepo.HasRejectedFriendRequestLog(frl)

	if err != nil {
		return err
	}
	if !canSend {
		return &e.BadRequestError{Message: "cannot send friend request: duplicate or active request exists"}
	}

	//make Action to UnFriended
	err = s.frlRepo.CreateFriendRequestLog(frl)

	if err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if err = s.fRepo.MakeUnFriend(user); err != nil {
		return err
	}

	if err = s.fRepo.MakeUnFriend(friend); err != nil {
		return err
	}

	return nil
}
