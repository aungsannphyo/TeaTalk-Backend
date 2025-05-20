package handler

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/success"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService s.UserService
}

func NewUserHandler(service s.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) RegisterHandler(c *gin.Context) {
	var user dto.RegisterRequestDto
	if err := c.ShouldBindJSON(&user); err != nil {
		e.BadRequestResponse(c, err)
		return
	}
	if err := dto.ValidateRegisterUser(user); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.userService.Register(&user); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.CreateResponse(c, "User have been created Successfull!")

}

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var u dto.LoginRequestDto

	if err := c.ShouldBindJSON(&u); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateLoginUser(u); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	foundUser, token, err := h.userService.Login(&u)

	if err != nil {
		e.UnauthorizedResponse(c, err)
		return
	}

	loginResponse := response.NewLoginResponse(foundUser, token)

	success.OkResponse(c, loginResponse)
}

func (h *UserHandler) GetUserHandler(c *gin.Context) {
	userId := c.Param("userID")

	user, err := h.userService.GetUserByID(c.Request.Context(), userId)
	if err != nil {
		e.NotFoundResponse(c, err)
		return
	}

	userResponse := response.NewUserResponse(user)

	success.OkResponse(c, userResponse)
}

func (h *UserHandler) GetChatListByUserIdHandler(c *gin.Context) {
	userID := c.GetString("userID")
	chatList, err := h.userService.GetChatListByUserID(c.Request.Context(), userID)

	if err != nil {
		e.NotFoundResponse(c, err)
		return
	}

	if len(chatList) == 0 {
		success.OkResponse(c, []models.ChatListItem{})
	} else {

		success.OkResponse(c, chatList)
	}

}

func (h *UserHandler) CreatePersonalDetailsHandler(c *gin.Context) {
	userID := c.GetString("userID")

	var pd dto.PersonalDetailDto
	if err := c.ShouldBindJSON(&pd); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateCreatePersonalDetails(pd); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.userService.CreatePersonalDetail(userID, &pd); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.CreateResponse(c, "Personal details created successfully!")
}

func (h *UserHandler) UpdatePersonalDetailsHandler(c *gin.Context) {
	userID := c.GetString("userID")
	var ps dto.PersonalDetailDto

	if err := c.ShouldBindJSON(&ps); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.userService.UpdatePersonalDetail(userID, &ps); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.CreateResponse(c, "Personal details have been updated Successfull!")
}

func (h *UserHandler) UploadProfileImageHandler(c *gin.Context) {
	userID := c.GetString("userID")

	file, err := c.FormFile("profile_image")
	if err != nil {
		e.BadRequestResponse(c, errors.New("image file is required"))
		return
	}

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("user_%s%s", userID, ext)
	savePath := filepath.Join("uploads", "profiles", filename)

	if err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	imageURL := "/static/profiles/" + filename

	// Update profile_image in DB
	if err := h.userService.UploadProfileImage(c.Request.Context(), userID, imageURL); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "Profile image updated successfully"})
}
