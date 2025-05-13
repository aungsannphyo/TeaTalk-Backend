package service

import "github.com/aungsannphyo/ywartalk/internal/domain/repository"

type FriendRequestLogService struct {
	frlRepo repository.FriendRequestLogRepository
}

func NewFriendRequestLogService(frlRepo repository.FriendRequestLogRepository) *FriendRequestLogService {
	return &FriendRequestLogService{
		frlRepo: frlRepo,
	}
}
