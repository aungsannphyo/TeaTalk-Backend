package services

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
)

type messageReadService struct {
	mrRepo repository.MessageReadRepository
}

func (s *messageReadService) CreateReadMessage(mr *models.MessageRead) error {
	err := s.mrRepo.CreateReadMessage(mr)

	if err != nil {
		return err
	}

	return nil
}

func (s *messageReadService) MarkAllReadMessages(userID, conversationID string) error {
	err := s.mrRepo.MarkAllReadMessages(userID, conversationID)

	if err != nil {
		return &e.InternalServerError{Message: "Failed to add sender to message read table"}
	}

	return nil
}
