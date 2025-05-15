package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type UserRepository interface {
	Register(user *models.User) error
	Login(user *models.User) (*models.User, error)
	GetUserById(userId string) (*models.User, error)
	GetGroupUsers(ctx context.Context, conversationId string) ([]models.User, error)
}
