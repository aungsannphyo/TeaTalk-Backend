package repository

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type MessageReadRepository interface {
	CreateReadMessage(mr *models.MessageRead) error
	MarkAllReadMessages(userID, conversationID string) error
}
