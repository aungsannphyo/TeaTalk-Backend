package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type ConversationRepository interface {
	CreateConversation(c *models.Conversation) error
	CheckExistsConversation(senderId, receiverId string) ([]models.Conversation, error)
}
