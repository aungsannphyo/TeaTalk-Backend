package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
)

type ConversationKeyService interface {
	CreateConversationKey(dto dto.CreateConversationKeyDto) error
	GetConversationKey(ctx context.Context, conversationID, userID string) (
		*response.ConversationKeyResponse,
		error,
	)
}
