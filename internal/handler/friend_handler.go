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
	var mufDto dto.UnFriendDto

	if err := c.ShouldBindJSON(&mufDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateUnFriendRequest(mufDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := h.fService.MakeUnFriend(mufDto, c); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{"message": "Successfully Unfriend!"})
}
