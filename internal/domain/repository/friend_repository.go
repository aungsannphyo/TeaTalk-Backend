package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type FriendRepository interface {
	CreateFriendShip(friend *models.Friend) error
	MakeUnFriend(FriendRepository *models.Friend) error
	AlreadyFriends(senderId, receiverId string) bool
}
