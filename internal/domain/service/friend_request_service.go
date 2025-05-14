package service

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type FriendRequestService interface {
	SendFriendRequest(dto dto.SendFriendRequestDto, c *gin.Context) error
	DecideFriendRequest(dto dto.DecideFriendRequestDto, c *gin.Context) error
}
