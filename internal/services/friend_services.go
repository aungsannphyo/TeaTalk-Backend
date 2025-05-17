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

	f := &models.Friend{
		UserID:   userID,
		FriendID: dto.FriendID,
	}

	frl := &models.FriendRequestLog{
		SenderID:    f.UserID,
		ReceiverID:  f.FriendID,
		Action:      models.ActionUnFriended,
		PerformedBy: f.UserID,
	}

	//make Action to UnFriended
	err := s.frlRepo.CreateFriendRequestLog(frl)

	if err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	return s.fRepo.MakeUnFriend(f)
}
