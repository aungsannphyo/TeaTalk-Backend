package service

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type FriendRequestService interface {
	SendFriendRequest(c *gin.Context, dto dto.SendFriendRequestDto) error
	DecideFriendRequest(c *gin.Context, dto dto.DecideFriendRequestDto) error
}
