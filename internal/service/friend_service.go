package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
)

type FriendService struct {
	fRepo   repository.FriendRepository
	frRepo  repository.FriendRequestRepository
	frlRepo repository.FriendRequestLogRepository
}

func NewFriendService(
	fRepo repository.FriendRepository,
	frRepo repository.FriendRequestRepository,
	frlRepo repository.FriendRequestLogRepository,
) *FriendService {
	return &FriendService{
		fRepo:   fRepo,
		frRepo:  frRepo,
		frlRepo: frlRepo,
	}
}

func (s *FriendService) MakeUnFriend(dto dto.UnFriendDto, c *gin.Context) error {

	f := &models.Friend{
		UserID:   c.GetString("userId"),
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
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	return s.fRepo.MakeUnFriend(f)
}
