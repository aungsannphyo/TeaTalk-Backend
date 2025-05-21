package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
)

type FriendRequestService interface {
	SendFriendRequest(ctx context.Context, userID string, dto dto.SendFriendRequestDto) error
	DecideFriendRequest(ctx context.Context, userID string, dto dto.DecideFriendRequestDto) error
	GetAllFriendRequestLog(ctx context.Context, userID string) ([]response.FriendRequestResponse, error)
}
