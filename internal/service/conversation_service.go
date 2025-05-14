package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConversationService struct {
	cRepo repository.ConversationRepository
}

func NewConversationService(cRepo repository.ConversationRepository) *ConversationService {
	return &ConversationService{
		cRepo: cRepo,
	}
}

func (s *ConversationService) CreateGroup(dto dto.CreateGroupDto, c *gin.Context) error {
	cID := uuid.NewString()
	userId := c.GetString("userId")

	conversation := &models.Conversation{
		ID:        cID,
		IsGroup:   true,
		Name:      &dto.Name,
		CreatedBy: &userId,
	}

	if err := s.cRepo.CreateConversation(conversation); err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}
	return nil
}
