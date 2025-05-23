package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type MessageReadService interface {
	CreateReadMessage(mr *models.MessageRead) error
}
