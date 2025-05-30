package handler

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

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

	foundUser, token, pdk, err := h.userService.Login(&u)

	if err != nil {
		e.UnauthorizedResponse(c, err)
		return
	}

	loginResponse := response.NewLoginResponse(foundUser, token, pdk)

	success.OkResponse(c, loginResponse)
}

func (h *UserHandler) GetChatListByUserIdHandler(c *gin.Context) {
	userID := c.GetString("userID")
	chatList, err := h.userService.GetChatListByUserID(c.Request.Context(), userID)

	if err != nil {
		e.NotFoundResponse(c, err)
		return
	}

	if len(chatList) == 0 {
		success.OkResponse(c, []response.ChatListResponse{})
	} else {

		success.OkResponse(c, chatList)
	}
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

func (h *UserHandler) SearchUserHandler(c *gin.Context) {
	input := c.Query("q")
	userID := c.GetString("userID")

	user, err := h.userService.SearchUser(c.Request.Context(), userID, input)

	if err != nil {
		e.NotFoundResponse(c, err)
		return
	}

	success.OkResponse(c, user)
}

func (h *UserHandler) GetUserHandler(c *gin.Context) {
	userID := c.Param("userID")

	user, ps, err := h.userService.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		e.NotFoundResponse(c, err)
		return
	}
	userDetailsResponse := response.NewUserDetailsResponse(user, ps)

	success.OkResponse(c, userDetailsResponse)
}

func (h *UserHandler) GetFriendsByUserHandler(c *gin.Context) {
	userID := c.GetString("userID")

	friend, err := h.userService.GetFriendsByID(c.Request.Context(), userID)
	if err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, friend)
}
