package services

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
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
