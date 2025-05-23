package handler

import (
	"strconv"
	"time"

	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/success"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	mService s.MessageService
}

func NewMessageHandler(s s.MessageService) *MessageHandler {
	return &MessageHandler{
		mService: s,
	}
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	conversationID := c.Param("conversationID")
	pageSizeStr := c.Query("pageSize")
	cursorTimeStr := c.Query("cursorTime")

	const defaultPageSize = 50
	pageSize, err := strconv.Atoi(pageSizeStr)

	if err != nil || pageSize <= 0 {
		pageSize = defaultPageSize
	}

	var cursorTime time.Time
	if cursorTimeStr != "" {
		cursorTime, err = time.Parse(time.RFC3339, cursorTimeStr)
		if err != nil {
			e.BadRequestResponse(c, err)
			return
		}
	} else {
		cursorTime = time.Now()
	}

	messages, err := h.mService.GetMessages(
		c.Request.Context(),
		conversationID,
		&cursorTime,
		pageSize,
	)

	if err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, messages)
}
