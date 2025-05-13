package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type UserRepository interface {
	Register(user *models.User) error
	Login(user *models.User) (*models.User, error)
	GetUser(userId string) (*models.User, error)
}
