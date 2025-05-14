package handler

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/service"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
)

type ConversationsHandler struct {
	cService *service.ConversationService
}

func NewConversationHandler(s *service.ConversationService) *ConversationsHandler {
	return &ConversationsHandler{
		cService: s,
	}
}

func (s *ConversationsHandler) CreateGroup(c *gin.Context) {
	var cgDto dto.CreateGroupDto

	if err := c.ShouldBindJSON(&cgDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateCreateGroup(cgDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.CreateGroup(cgDto, c); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.CreateResponse(c, "You have been successfully created the group!")
}
