package handler

import (
	"fmt"

	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
)

type FriendRequestHandler struct {
	frService s.FriendRequestService
}

func NewFriendRequestHandler(service s.FriendRequestService) *FriendRequestHandler {
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

	if err := h.frService.SendFriendRequest(frDto, c); err != nil {
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

	if err := h.frService.DecideFriendRequest(dfrDto, c); err != nil {
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
		"message": fmt.Sprintf("You have successfully %v for a friend request", dfrDto.Status),
	})

}
