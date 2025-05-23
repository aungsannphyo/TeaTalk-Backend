package handler

import (
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/success"
	"github.com/gin-gonic/gin"
)

type MessageReadHandler struct {
	mrService s.MessageReadService
}

func NewMessageReadHandler(s s.MessageReadService) *MessageReadHandler {
	return &MessageReadHandler{
		mrService: s,
	}
}

func (h *MessageReadHandler) MarkAllReadMessagesHandler(c *gin.Context) {
	userID := c.GetString("userID")
	conversationID := c.Param("conversationID")

	err := h.mrService.MarkAllReadMessages(userID, conversationID)

	if err != nil {
		e.InternalServerResponse(c, err)
	}

	success.CreateResponse(c, "All Messges have been read")
}
