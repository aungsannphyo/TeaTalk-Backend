package handler

import (
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

	user, err := h.userService.GetUserById(c.Request.Context(), userId)
	if err != nil {
		e.NotFoundResponse(c, err)
		return
	}

	userResponse := response.NewUserResponse(user)

	success.OkResponse(c, userResponse)
}

func (h *UserHandler) GetChatListByUserId(c *gin.Context) {
	userID := c.GetString("userID")
	chatList, err := h.userService.GetChatListByUserId(c.Request.Context(), userID)

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
