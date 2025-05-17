package handler

import (
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/success"
	"github.com/gin-gonic/gin"
)

type FriendHandler struct {
	fService s.FriendService
}

func NewFriendHandler(s s.FriendService) *FriendHandler {
	return &FriendHandler{
		fService: s,
	}
}

func (h *FriendHandler) MakeUnFriendHandler(c *gin.Context) {
	var mufDto dto.UnFriendDto
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&mufDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateUnFriendRequest(mufDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.fService.MakeUnFriend(userID, mufDto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "Successfully Unfriend!"})
}
