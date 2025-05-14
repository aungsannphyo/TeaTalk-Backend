package service

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type FriendService interface {
	MakeUnFriend(dto dto.UnFriendDto, c *gin.Context) error
}
