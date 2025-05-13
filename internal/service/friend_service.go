package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type FriendService struct {
	fRepo repository.FriendRepository
}

func NewFriendService(fRepo repository.FriendRepository) *FriendService {
	return &FriendService{
		fRepo: fRepo,
	}
}

func (s *FriendService) CreateFriendShip(fDto *dto.CreateFriendDto) error {
	friend := &models.Friend{
		UserID:   fDto.UserID,
		FriendID: fDto.FriendID,
	}
	return s.fRepo.CreateFriendShip(friend)
}

func (s *FriendService) MakeUnFriend(ufDto *dto.UnFriendDto) error {
	uf := &models.Friend{
		UserID:   ufDto.UserID,
		FriendID: ufDto.FriendID,
	}
	return s.fRepo.MakeUnFriend(uf)
}
