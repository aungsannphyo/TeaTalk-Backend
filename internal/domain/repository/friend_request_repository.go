package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type FriendRequestRepository interface {
	SendFriendRequest(fr *models.FriendRequest) error
	RejectFriendRequest(fr *models.FriendRequest) error
	HasPendingRequest(senderId, receiverId string) bool
	DeleteById(id string) error
	FindById(id string) (*models.FriendRequest, error)
}
