package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type UserService interface {
	Register(u *dto.RegisterRequestDto) error
	Login(u *dto.LoginRequestDto) (*models.User, string, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	GetChatListByUserID(ctx context.Context, userID string) ([]models.ChatListItem, error)
	CreatePersonalDetail(userID string, pd *dto.PersonalDetailDto) error
	UpdatePersonalDetail(userID string, pd *dto.PersonalDetailDto) error
	UploadProfileImage(ctx context.Context, userID string, imagePath string) error
}
