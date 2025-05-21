package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type FriendRequestRepository interface {
	SendFriendRequest(fr *models.FriendRequest) error
	RejectFriendRequest(fr *models.FriendRequest) error
	HasPendingRequest(ctx context.Context, senderId, receiverId string) bool
	DeleteFriendRequestByID(id string) error
	GetFriendRequestByID(ctx context.Context, id string) (*models.FriendRequest, error)
}
