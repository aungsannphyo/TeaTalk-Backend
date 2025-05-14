package handler

import (
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	"github.com/aungsannphyo/ywartalk/pkg/common"
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
		common.BadRequestResponse(c, err)
		return
	}
	if err := dto.ValidateRegisterUser(user); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := h.userService.Register(&user); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.CreateResponse(c, "User have been created Successfull!")

}

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var u dto.LoginRequestDto

	if err := c.ShouldBindJSON(&u); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateLoginUser(u); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	foundUser, token, err := h.userService.Login(&u)

	if err != nil {
		common.UnauthorizedResponse(c, err)
		return
	}

	loginResponse := response.NewLoginResponse(foundUser, token)

	common.OkResponse(c, loginResponse)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userId := c.Param("id")

	user, err := h.userService.GetUser(userId)
	if err != nil {
		common.NotFoundResponse(c, err)
		return
	}

	userResponse := response.NewUserResponse(user)

	common.OkResponse(c, userResponse)
}
