package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type ConversationRepository interface {
	CreateConversation(c *models.Conversation) error
	CheckExistsConversation(senderId, receiverId string) ([]models.Conversation, error)
	UpdateGroupName(c *models.Conversation) error
	InviteGroup(c *models.Conversation) error
}
