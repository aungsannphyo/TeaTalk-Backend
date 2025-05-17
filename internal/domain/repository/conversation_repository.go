package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type ConversationRepository interface {
	CreateConversation(c *models.Conversation) error
	CheckExistsConversation(ctx context.Context, senderId, receiverId string) (models.Conversation, error)
	UpdateGroupName(c *models.Conversation) error
	CheckExistsGroup(ctx context.Context, c *models.Conversation) bool
	GetGroupMembers(ctx context.Context, conversationId string) ([]models.User, error)
	GetGroupsById(ctx context.Context, userID string) ([]models.Conversation, error)
}
