package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
)

type UserRepository interface {
	Register(user *models.User) error
	Login(user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, *models.PersonalDetails, error)
	GetChatListByUserID(ctx context.Context, userID string) ([]response.ChatListResponse, error)
	UpdatePersonalDetail(pd *models.PersonalDetails) error
	GetProfileImagePath(ctx context.Context, userID string) (string, error)
	UploadProfileImage(userID string, imagePath string) error
	SearchUser(ctx context.Context, userID string, searchInput string) ([]response.SearchResultResponse, error)
	SetUserOnline(userID string) error
	SetUserOffline(userID string) error
	GetFriendsByID(ctx context.Context, userID string) ([]response.FriendResponse, error)
	GetUserKeyByID(ctx context.Context, userID string) (*response.UserKeyResponse, error)
}
