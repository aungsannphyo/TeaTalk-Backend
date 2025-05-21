package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
)

type UserRepository interface {
	Register(user *models.User) error
	Login(user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	GetChatListByUserID(ctx context.Context, userID string) ([]models.ChatListItem, error)
	CreatePersonalDetail(ps *models.PersonalDetails) error
	UpdatePersonalDetail(pd *models.PersonalDetails) error
	GetProfileImagePath(ctx context.Context, userID string) (string, error)
	UploadProfileImage(userID string, imagePath string) error
	SearchUser(ctx context.Context, userID string, searchInput string) ([]response.SearchResultResponse, error)
}
