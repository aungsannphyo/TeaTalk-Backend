package handler

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/service"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
)

type FriendHandler struct {
	fService *service.FriendService
}

func NewFriendHandler(s *service.FriendService) *FriendHandler {
	return &FriendHandler{
		fService: s,
	}
}

func (h *FriendHandler) MakeUnFriend(c *gin.Context) {
	var duf dto.UnFriendDto

	if err := c.ShouldBindJSON(&duf); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateUnFriendRequest(duf); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := h.fService.MakeUnFriend(&duf); err != nil {
		common.InternalServerResponse(c, err)
	}

	common.OkResponse(c, gin.H{"message": "Successfully Unfriend!"})
}
