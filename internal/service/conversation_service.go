package service

import "github.com/aungsannphyo/ywartalk/internal/domain/repository"

type ConversationService struct {
	cRepo repository.ConversationRepository
}

func NewConversationService(cRepo repository.ConversationRepository) *ConversationService {
	return &ConversationService{
		cRepo: cRepo,
	}
}
