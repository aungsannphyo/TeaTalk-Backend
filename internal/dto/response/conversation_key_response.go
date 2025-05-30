package response

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
)

type ConversationKeyResponse struct {
	ConversationID           string `json:"conversationId"`
	UserID                   string `json:"userId"`
	ConversationEncryptedKey string `json:"conversationEncryptedKey"`
	ConversationKeyNonce     string `json:"conversationKeyNonce"`
}

func NewConversationKeyResponse(c *models.ConversationKey) *ConversationKeyResponse {
	return &ConversationKeyResponse{
		ConversationID:           c.ConversationId,
		UserID:                   c.UserID,
		ConversationEncryptedKey: utils.EncodeBase64(c.ConversationEncryptedKey),
		ConversationKeyNonce:     utils.EncodeBase64(c.ConversationKeyNonce),
	}
}
