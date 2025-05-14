package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type MessageRepository interface {
	CreateMessage(m *models.Message) error
}
