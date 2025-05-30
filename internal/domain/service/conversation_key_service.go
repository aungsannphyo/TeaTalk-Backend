package service

import "github.com/aungsannphyo/ywartalk/internal/dto"

type ConversationKeyService interface {
	CreateConversationKey(dto dto.CreateConversationKeyDto) error
}
