package handler

import (
	"fmt"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
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
	var frDto dto.SendFriendRequestDto

	if err := c.ShouldBindJSON(&frDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateSendFriendRequest(frDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	fr := &models.FriendRequest{
		SenderId:   c.GetString("userId"),
		ReceiverId: frDto.ReceiverId,
	}

	if err := h.frService.SendFriendRequest(fr); err != nil {
		common.ConfictResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{"message": "You have successfully sent a friend request"})
}

func (h *FriendRequestHandler) DecideFriendRequest(c *gin.Context) {
	var dfrDto dto.DecideFriendRequestDto

	if err := c.ShouldBindJSON(&dfrDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateDecideFriendRequest(dfrDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	dfr := &models.FriendRequest{
		ID:         dfrDto.FriendRequestId,
		ReceiverId: c.GetString("userId"),
		Status:     dfrDto.Status,
	}

	if err := h.frService.DecideFriendRequest(dfr); err != nil {
		if _, ok := err.(*common.ForbiddenError); ok {
			common.ConfictResponse(c, err)
			return
		}

		if _, ok := err.(*common.InternalServerError); ok {
			common.InternalServerResponse(c, err)
			return
		}

		if _, ok := err.(*common.NotFoundError); ok {
			common.NotFoundResponse(c, err)
			return
		}

	}
	common.OkResponse(c, gin.H{
		"message": fmt.Sprintf("You have successfully %v for a friend request", dfr.Status),
	})

}
