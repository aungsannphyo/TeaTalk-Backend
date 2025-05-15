package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type FriendRepository interface {
	CreateFriendShip(friend *models.Friend) error
	MakeUnFriend(FriendRepository *models.Friend) error
	AlreadyFriends(ctx context.Context, senderId, receiverId string) bool
}
