package services

import (
	"context"
	"errors"
	"log"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
	"github.com/go-sql-driver/mysql"
)

type userServices struct {
	userRepo repository.UserRepository
}

func (s *userServices) Register(u *dto.RegisterRequestDto) error {
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return &e.InternalServerError{Message: "Password hashing failed"}
	}
	user := &models.User{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
	}

	err = s.userRepo.Register(user)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return &e.BadRequestError{Message: "Email already exists"}
			}
		}
		log.Println("ERR", err)
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	return nil
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
		return nil, "", &e.UnAuthorizedError{Message: "Password doesn't match"}
	}

	token, err := utils.GenerateToken(foundUser.Email, foundUser.ID)

	if err != nil {
		return nil, "", &e.InternalServerError{Message: "Something went wrong while generating token"}
	}

	return foundUser, token, nil
}

func (s *userServices) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	user, err := s.userRepo.GetUserById(ctx, userId)

	if err != nil {
		return nil, &e.NotFoundError{Message: "User not found"}

	}
	return user, nil
}

func (s *userServices) GetChatListByUserId(ctx context.Context, userID string) ([]models.ChatListItem, error) {
	chatList, err := s.userRepo.GetChatListByUserId(ctx, userID)

	if err != nil {
		return nil, &e.InternalServerError{Message: err.Error()}
	}

	return chatList, nil
}
