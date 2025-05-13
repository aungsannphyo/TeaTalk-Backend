package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(u *dto.RegisterRequestDto) error {
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	user := &models.User{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
	}
	return s.userRepo.Register(user)
}

func (s *UserService) Login(u *dto.LoginRequestDto) (*models.User, string, error) {
	user := &models.User{
		Email:    u.Email,
		Password: u.Password,
	}

	foundUser, err := s.userRepo.Login(user)

	if err != nil {
		return nil, "", err
	}

	checkPassword := utils.CheckPasswordHash(user.Password, foundUser.Password)

	if !checkPassword {
		return nil, "", err
	}

	token, err := utils.GenerateToken(foundUser.Email, foundUser.ID)

	if err != nil {
		return nil, "", err
	}

	return foundUser, token, nil
}

func (s *UserService) GetUser(userId string) (*models.User, error) {
	user, err := s.userRepo.GetUser(userId)

	if err != nil {
		return nil, err

	}
	return user, nil
}
