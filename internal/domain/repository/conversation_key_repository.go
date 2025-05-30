package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type ConversationKeyRepository interface {
	CreateConversationKey(cKeys *models.ConversationKey) error
	GetConversationKey(ctx context.Context, conversationID, userID string) (*models.ConversationKey, error)
}
