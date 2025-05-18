package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type UserService interface {
	Register(u *dto.RegisterRequestDto) error
	Login(u *dto.LoginRequestDto) (*models.User, string, error)
	GetUserById(ctx context.Context, userID string) (*models.User, error)
	GetChatListByUserId(ctx context.Context, userID string) ([]models.ChatListItem, error)
}
