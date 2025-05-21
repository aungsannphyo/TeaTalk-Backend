package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
)

type FriendRequestLogRepository interface {
	CreateFriendRequestLog(frl *models.FriendRequestLog) error
	GetAllFriendRequestLog(ctx context.Context, userID string) ([]response.FriendRequestResponse, error)
}
