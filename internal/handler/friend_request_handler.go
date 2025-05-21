package handler

import (
	"fmt"

	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/success"
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

func (h *FriendRequestHandler) SendFriendRequestHandler(c *gin.Context) {
	var frDto dto.SendFriendRequestDto
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&frDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateSendFriendRequest(frDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.frService.SendFriendRequest(c.Request.Context(), userID, frDto); err != nil {
		e.ConflictResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have successfully sent a friend request"})
}

func (h *FriendRequestHandler) DecideFriendRequestHandler(c *gin.Context) {
	var dfrDto dto.DecideFriendRequestDto
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&dfrDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateDecideFriendRequest(dfrDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.frService.DecideFriendRequest(c.Request.Context(), userID, dfrDto); err != nil {
		if _, ok := err.(*e.ForbiddenError); ok {
			e.ConflictResponse(c, err)
			return
		}

		if _, ok := err.(*e.InternalServerError); ok {
			e.InternalServerResponse(c, err)
			return
		}

		if _, ok := err.(*e.NotFoundError); ok {
			e.NotFoundResponse(c, err)
			return
		}

	}
	success.OkResponse(c, gin.H{
		"message": fmt.Sprintf("You have successfully %v for a friend request", dfrDto.Status),
	})
}

func (h *FriendRequestHandler) GetAllFriendRequestHandler(c *gin.Context) {
	userID := c.GetString("userID")

	logs, err := h.frService.GetAllFriendRequestLog(c.Request.Context(), userID)

	if err != nil {
		e.InternalServerResponse(c, err)
	}

	success.OkResponse(c, logs)
}
