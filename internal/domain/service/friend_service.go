package service

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type FriendService interface {
	MakeUnFriend(userID string, dto dto.UnFriendDto) error
}
