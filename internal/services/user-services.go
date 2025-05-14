package services

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
)

type userServices struct {
	userRepo repository.UserRepository
}

func (s *userServices) Register(u *dto.RegisterRequestDto) error {
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}
	user := &models.User{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
	}
	return s.userRepo.Register(user)
}

func (s *userServices) Login(u *dto.LoginRequestDto) (*models.User, string, error) {
	user := &models.User{
		Email:    u.Email,
		Password: u.Password,
	}

	foundUser, err := s.userRepo.Login(user)

	if err != nil {
		return nil, "", &common.UnAuthorizedError{Message: "Email or Password doesn't match"}
	}

	checkPassword := utils.CheckPasswordHash(user.Password, foundUser.Password)

	if !checkPassword {
		return nil, "", &common.InternalServerError{Message: "Something went wrong, Please try agian later"}
	}

	token, err := utils.GenerateToken(foundUser.Email, foundUser.ID)

	if err != nil {
		return nil, "", &common.InternalServerError{Message: "Something went wrong, Please try agian later"}
	}

	return foundUser, token, nil
}

func (s *userServices) GetUser(userId string) (*models.User, error) {
	user, err := s.userRepo.GetUser(userId)

	if err != nil {
		return nil, &common.NotFoundError{Message: "User not found"}

	}
	return user, nil
}
