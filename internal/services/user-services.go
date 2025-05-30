package services

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type userServices struct {
	userRepo repository.UserRepository
}

func (s *userServices) Register(u *dto.RegisterRequestDto) error {

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return &e.InternalServerError{Message: "Password hashing failed"}
	}

	salt, saleErr := utils.DecodeBase64(u.Salt)
	encryptedUserKey, keyErr := utils.DecodeBase64(u.EncryptedUserKey)
	userKeyNonce, nonceErr := utils.DecodeBase64(u.UserKeyNonce)

	if saleErr != nil || keyErr != nil || nonceErr != nil {
		return &e.InternalServerError{Message: "Fail to decode string"}
	}

	user := &models.User{
		ID:               uuid.New().String(),
		Username:         u.Username,
		Email:            u.Email,
		Password:         hashedPassword,
		Salt:             salt,
		EncryptedUserKey: encryptedUserKey,
		UserKeyNonce:     userKeyNonce,
	}

	err = s.userRepo.Register(user)
	if err != nil {
		log.Println("ERR", err)
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return &e.BadRequestError{Message: "Email already exists"}
			}
		}
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	return nil
}

func (s *userServices) Login(u *dto.LoginRequestDto) (*models.User, string, []byte, error) {
	user := &models.User{
		Email:    u.Email,
		Password: u.Password,
	}

	foundUser, err := s.userRepo.Login(user)

	if err != nil {
		return nil, "", nil, &e.UnAuthorizedError{Message: "Email or Password doesn't match"}
	}

	checkPassword := utils.CheckPasswordHash(user.Password, foundUser.Password)

	if !checkPassword {
		return nil, "", nil, &e.UnAuthorizedError{Message: "Password doesn't match"}
	}

	if err != nil {
		return nil, "", nil, &e.InternalServerError{Message: "Something went wrong while decryption key"}
	}

	token, err := utils.GenerateToken(foundUser.Email, foundUser.ID)

	if err != nil {
		return nil, "", nil, &e.InternalServerError{Message: "Something went wrong while generating token"}
	}

	pdk := utils.DeriveKey([]byte(u.Password), foundUser.Salt)

	return foundUser, token, pdk, nil
}

func (s *userServices) GetUserByID(ctx context.Context, userID string) (*models.User, *models.PersonalDetails, error) {
	user, ps, err := s.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return nil, nil, &e.NotFoundError{Message: "User not found"}

	}
	return user, ps, nil
}

func (s *userServices) GetChatListByUserID(ctx context.Context, userID string) ([]response.ChatListResponse, error) {
	chatList, err := s.userRepo.GetChatListByUserID(ctx, userID)
	if err != nil {
		return nil, &e.InternalServerError{Message: err.Error()}
	}

	var list []response.ChatListResponse
	for _, chat := range chatList {
		l := response.ChatListResponse{
			ConversationID:       chat.ConversationID,
			IsGroup:              chat.IsGroup,
			Name:                 chat.Name,
			LastMessageID:        chat.LastMessageID,
			LastMessageContent:   chat.LastMessageContent,
			LastMessageSender:    chat.LastMessageSender,
			LastMessageCreatedAt: chat.LastMessageCreatedAt,
			UnreadCount:          chat.UnreadCount,
			ReceiverID:           chat.ReceiverID,
			ProfileImage:         chat.ProfileImage,
			TotalOnline:          chat.TotalOnline,
			LastSeen:             chat.LastSeen,
		}
		list = append(list, l)
	}

	return list, nil
}

func (s *userServices) UpdatePersonalDetail(userID string, dto *dto.PersonalDetailDto) error {
	ps := &models.PersonalDetails{
		UserID:      userID,
		Gender:      dto.Gender,
		DateOfBirth: dto.DateOfBirth,
		Bio:         dto.Bio,
	}

	err := s.userRepo.UpdatePersonalDetail(ps)

	if err != nil {
		return &e.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *userServices) UploadProfileImage(ctx context.Context, userID string, imagePath string) error {
	oldPath, err := s.userRepo.GetProfileImagePath(ctx, userID)

	if err != nil {
		return &e.InternalServerError{Message: err.Error()}
	}

	if oldPath != "" {
		filePath := "." + oldPath
		os.Remove(filePath)
		return nil
	} else {
		err := s.userRepo.UploadProfileImage(userID, imagePath)

		if err != nil {
			return &e.InternalServerError{Message: err.Error()}
		}
	}

	err = s.userRepo.UploadProfileImage(userID, imagePath)

	if err != nil {
		return &e.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *userServices) SearchUser(ctx context.Context, userID string, searchInput string) ([]response.SearchResultResponse, error) {
	users, err := s.userRepo.SearchUser(ctx, userID, searchInput)

	if err != nil {
		return nil, &e.NotFoundError{Message: "User not found"}
	}

	return users, nil
}

func (s *userServices) SetUserOnline(userID string) error {
	return s.userRepo.SetUserOnline(userID)
}

func (s *userServices) SetUserOffline(userID string) error {
	return s.userRepo.SetUserOffline(userID)
}

func (s *userServices) GetFriendsByID(ctx context.Context, userID string) (
	[]response.FriendResponse,
	error,
) {
	friend, err := s.userRepo.GetFriendsByID(ctx, userID)

	if err != nil {
		return nil, &e.InternalServerError{Message: "Something went wrong.Please try again later!"}

	}
	return friend, nil
}
