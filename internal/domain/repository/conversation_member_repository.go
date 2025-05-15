package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type ConversationMemeberRepository interface {
	CreateConversationMember(cm *models.ConversationMember) error
	CheckConversationMember(ctx context.Context, cm *models.ConversationMember) bool
}
