package service

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type MessageService interface {
	SendPrivateMessage(c *gin.Context, dto dto.SendPrivateMessageDto) error
	SendGroupMessage(c *gin.Context, dto dto.SendGroupMessageDto) error
}
