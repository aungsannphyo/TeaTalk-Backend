package handler

import (
	"fmt"

	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/service"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
)

type FriendRequestHandler struct {
	frService *service.FriendRequestService
}

func NewFriendRequestHandler(service *service.FriendRequestService) *FriendRequestHandler {
	return &FriendRequestHandler{
		frService: service,
	}
}

func (h *FriendRequestHandler) SendFriendRequest(c *gin.Context) {
	var fr dto.SendFriendRequestDto

	if err := c.ShouldBindJSON(&fr); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateSendFriendRequest(fr); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := h.frService.SendFriendRequest(&fr); err != nil {
		common.ConfictResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{"message": "You have successfully sent a friend request"})
}

func (h *FriendRequestHandler) DecideFriendRequest(c *gin.Context) {
	var dfr dto.DecideFriendRequestDto

	if err := c.ShouldBindJSON(&dfr); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateDecideFriendRequest(dfr); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := h.frService.DecideFriendRequest(&dfr); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{
		"message": fmt.Sprintf("You have successfully %v for a friend request", dfr.Status),
	})

}
