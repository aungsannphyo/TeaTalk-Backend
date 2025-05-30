package handler

import (
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/success"
	"github.com/gin-gonic/gin"
)

type ConversationKeyHandler struct {
	cKeyService s.ConversationKeyService
}

func NewConversationKeyHandler(s s.ConversationKeyService) *ConversationKeyHandler {
	return &ConversationKeyHandler{
		cKeyService: s,
	}
}

func (h *ConversationKeyHandler) CreateConversationKey(c *gin.Context) {
	var ckDto dto.CreateConversationKeyDto

	if err := c.ShouldBindJSON(&ckDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}
	if err := dto.ValidateCreateConversationKey(ckDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.cKeyService.CreateConversationKey(ckDto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.CreateResponse(c, "You have been successfully created the conversation key!")

}
