package response

import (
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
)

type LoginResponse struct {
	ID               string    `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Token            string    `json:"token"`
	PDK              string    `json:"pdk"`
	EncryptedUserKey string    `json:"encryptedUserKey"`
	UserKeyNonce     string    `json:"userKeyNonce"`
	CreatedAt        time.Time `json:"createdAt"`
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

func NewLoginResponse(user *models.User, token string, pdk []byte) *LoginResponse {
	return &LoginResponse{
		ID:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		CreatedAt:        user.CreatedAt,
		PDK:              utils.EncodeBase64(pdk),
		EncryptedUserKey: utils.EncodeBase64(user.EncryptedUserKey),
		UserKeyNonce:     utils.EncodeBase64(user.UserKeyNonce),
		Token:            token,
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
