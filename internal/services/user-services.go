package services

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
	"github.com/gin-gonic/gin"
)

type userServices struct {
	userRepo repository.UserRepository
}

func (s *userServices) Register(u *dto.RegisterRequestDto) error {
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
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
		return nil, "", &e.UnAuthorizedError{Message: "Email or Password doesn't match"}
	}

	checkPassword := utils.CheckPasswordHash(user.Password, foundUser.Password)

	if !checkPassword {
		return nil, "", &e.InternalServerError{Message: "Something went wrong, Please try agian later"}
	}

	token, err := utils.GenerateToken(foundUser.Email, foundUser.ID)

	if err != nil {
		return nil, "", &e.InternalServerError{Message: "Something went wrong, Please try agian later"}
	}

	return foundUser, token, nil
}

func (s *userServices) GetUserById(c *gin.Context, userId string) (*models.User, error) {
	user, err := s.userRepo.GetUserById(c, userId)

	if err != nil {
		return nil, &e.NotFoundError{Message: "User not found"}

	}
	return user, nil
}

func (s *userServices) GetGroupUsers(c *gin.Context, conversationId string) ([]models.User, error) {
	users, err := s.userRepo.GetGroupUsers(c.Request.Context(), conversationId)
	if err != nil {
		return nil, &e.InternalServerError{Message: err.Error()}
	}
	return users, nil
}
