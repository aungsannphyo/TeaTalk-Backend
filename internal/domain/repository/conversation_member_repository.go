package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type ConversationMemeberRepository interface {
	CreateConversationMember(cm *models.ConversationMember) error
	CheckConversationMember(cm *models.ConversationMember) bool
}
