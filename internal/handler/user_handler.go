package handler

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
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

func (h *UserHandler) GetUserById(c *gin.Context) {
	userId := c.Param("id")

	user, err := h.userService.GetUserById(c, userId)
	if err != nil {
		common.NotFoundResponse(c, err)
		return
	}

	userResponse := response.NewUserResponse(user)

	common.OkResponse(c, userResponse)
}

func (h *UserHandler) GetGroupUsers(c *gin.Context) {
	groupId := c.Param("groupId")

	groupUsers, err := h.userService.GetGroupUsers(c, groupId)

	if err != nil {
		common.InternalServerResponse(c, err)
	}

	var users []response.UserResponse

	for _, groupUser := range groupUsers {
		user := &models.User{
			ID:        groupUser.ID,
			Email:     groupUser.Email,
			Username:  groupUser.Username,
			CreatedAt: groupUser.CreatedAt,
		}
		userResponse := response.NewUserResponse(user)
		users = append(users, *userResponse)
	}

	common.OkResponse(c, users)
}
