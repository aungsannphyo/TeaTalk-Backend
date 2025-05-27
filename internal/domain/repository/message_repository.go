package repository

import (
	"context"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type MessageRepository interface {
	CreateMessage(m *models.Message) error
	GetMessages(
		ctx context.Context,
		conversationID string,
		cursorTimestamp *time.Time,
		pageSize int,
	) ([]dto.MessagesDto, error)
}
