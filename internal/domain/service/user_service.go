package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(u *dto.RegisterRequestDto) error
	Login(u *dto.LoginRequestDto) (*models.User, string, error)
	GetUserById(c *gin.Context, userId string) (*models.User, error)
	GetGroupUsers(c *gin.Context, conversationId string) ([]models.User, error)
}
