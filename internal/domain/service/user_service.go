package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
)

type UserService interface {
	Register(u *dto.RegisterRequestDto) error
	Login(u *dto.LoginRequestDto) (*models.User, string, []byte, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, *models.PersonalDetails, error)
	GetChatListByUserID(ctx context.Context, userID string) ([]response.ChatListResponse, error)
	UpdatePersonalDetail(userID string, pd *dto.PersonalDetailDto) error
	UploadProfileImage(ctx context.Context, userID string, imagePath string) error
	SearchUser(ctx context.Context, userID string, searchInput string) ([]response.SearchResultResponse, error)
	SetUserOnline(userID string) error
	SetUserOffline(userID string) error
	GetFriendsByID(ctx context.Context, userID string) ([]response.FriendResponse, error)
	UpdateUserName(userID string, username string) error
}
