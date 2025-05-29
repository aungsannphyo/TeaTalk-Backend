package response

import (
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type LoginResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserResponse struct {
	ID           string    `json:"id"`
	UserIdentity string    `json:"userIdentity"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"createdAt"`
}

type SearchResultResponse struct {
	ID           string     `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	UserIdentity string     `json:"userIdentity"`
	IsFriend     bool       `json:"isFriend"`
	ProfileImage *string    `json:"profileImage"`
	IsOnline     bool       `json:"isOnline"`
	LastSeen     *time.Time `json:"lastSeen"`
}

type UserDetailsResponse struct {
	User    UserResponse            `json:"user"`
	Details PersonalDetailsResponse `json:"personalDetails"`
}

type UserKeyResponse struct {
	UserID           string `json:"id"`
	Salt             []byte `json:"salt"`
	EncryptedUserKey []byte `json:"encryptedUserKey"`
	UserKeyNonce     []byte `json:"userKeyNonce"`
}

func NewLoginResponse(user *models.User, token string) *LoginResponse {
	return &LoginResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}

}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:           user.ID,
		UserIdentity: user.UserIdentity,
		Username:     user.Username,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
	}
}

func NewUserDetailsResponse(user *models.User, details *models.PersonalDetails) *UserDetailsResponse {
	return &UserDetailsResponse{
		User:    *NewUserResponse(user),
		Details: *NewPersonalDetailsResponse(details),
	}
}
