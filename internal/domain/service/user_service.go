package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(u *dto.RegisterRequestDto) error
	Login(u *dto.LoginRequestDto) (*models.User, string, error)
	GetUserById(userId string) (*models.User, error)
	GetGroupUsers(conversationId string, c *gin.Context) ([]models.User, error)
}
