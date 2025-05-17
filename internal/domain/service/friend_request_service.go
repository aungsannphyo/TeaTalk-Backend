package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type FriendRequestService interface {
	SendFriendRequest(ctx context.Context, userID string, dto dto.SendFriendRequestDto) error
	DecideFriendRequest(ctx context.Context, userID string, dto dto.DecideFriendRequestDto) error
}
