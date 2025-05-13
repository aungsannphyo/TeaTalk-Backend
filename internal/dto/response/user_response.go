package response

import (
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type loginResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type userResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewLoginResponse(user *models.User, token string) *loginResponse {
	return &loginResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}

}

func NewUserResponse(user *models.User) *userResponse {
	return &userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
