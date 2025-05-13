package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type FriendRequestRepository interface {
	SendFriendRequest(fr *models.FriendRequest) error
	DecideFriendRequest(fr *models.FriendRequest) error
	AlreadyFriends(senderId, receiverId string) (bool, error)
	HasPendingRequest(senderId, receiverId string) (bool, error)
	FindById(friendRequestId string) (*models.FriendRequest, error)
}
