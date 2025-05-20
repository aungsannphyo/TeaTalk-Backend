package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type UserRepository interface {
	Register(user *models.User) error
	Login(user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, userID string) (*models.User, error)
	GetChatListByUserId(ctx context.Context, userID string) ([]models.ChatListItem, error)
	CreatePersonalDetail(ps *models.PersonalDetails) error
}
