package service

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type MessageService interface {
	SendPrivateMessage(dto dto.SendPrivateMessageDto, c *gin.Context) error
}
