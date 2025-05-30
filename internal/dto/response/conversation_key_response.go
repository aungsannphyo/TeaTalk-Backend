package response

import (
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type ConversationKeyResponse struct {
	ConversationID           string    `json:"conversationId"`
	UserID                   string    `json:"userId"`
	ConversationEncryptedKey []byte    `json:"conversationEncryptedKey"`
	ConversationKeyNonce     []byte    `json:"conversationKeyNonce"`
	CreatedAt                time.Time `json:"createdAt"`
}

func NewConversationKeyResponse(c *models.ConversationKey) *ConversationKeyResponse {
	return &ConversationKeyResponse{
		ConversationID:           c.ConversationId,
		UserID:                   c.UserID,
		ConversationEncryptedKey: c.ConversationEncryptedKey,
		ConversationKeyNonce:     c.ConversationKeyNonce,
	}
}
