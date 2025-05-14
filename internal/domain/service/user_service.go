package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type UserService interface {
	Register(u *dto.RegisterRequestDto) error
	Login(u *dto.LoginRequestDto) (*models.User, string, error)
	GetUser(userId string) (*models.User, error)
}
