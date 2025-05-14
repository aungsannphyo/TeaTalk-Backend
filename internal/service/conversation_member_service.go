package service

import "github.com/aungsannphyo/ywartalk/internal/domain/repository"

type ConversationMemberService struct {
	cmRepo repository.ConversationMemeberRepository
}

func NewConversationMemberService(cmRepo repository.ConversationMemeberRepository) *ConversationMemberService {
	return &ConversationMemberService{
		cmRepo: cmRepo,
	}
}
