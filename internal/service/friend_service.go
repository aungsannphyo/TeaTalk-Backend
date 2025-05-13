package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
)

type FriendService struct {
	fRepo  repository.FriendRepository
	frRepo repository.FriendRequestRepository
}

func NewFriendService(fRepo repository.FriendRepository,
	frRepo repository.FriendRequestRepository) *FriendService {
	return &FriendService{
		fRepo:  fRepo,
		frRepo: frRepo,
	}
}

func (s *FriendService) CreateFriendShip(f *models.Friend) error {
	return s.fRepo.CreateFriendShip(f)
}

func (s *FriendService) MakeUnFriend(f *models.Friend) error {
	return s.fRepo.MakeUnFriend(f)
}
