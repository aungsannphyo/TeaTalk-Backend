package repository

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type FriendRequestLogRepository interface {
	CreateFriendRequestLog(frl *models.FriendRequestLog) error
}
