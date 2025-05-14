package handler

import (
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	mService s.MessageService
}

func NewMessageHandler(service s.MessageService) *MessageHandler {
	return &MessageHandler{
		mService: service,
	}
}

func (h *MessageHandler) SendPrivateMessage(c *gin.Context) {

	var smDto dto.SendPrivateMessageDto

	if err := c.ShouldBindJSON(&smDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateSendMessageRequest(smDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := h.mService.SendPrivateMessage(smDto, c); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{"message": "Message send successfully!"})
}
