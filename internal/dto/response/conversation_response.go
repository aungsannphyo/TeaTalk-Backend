package response

import (
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type ConversationResponse struct {
	ConversationID string    `json:"conversationId"`
	IsGroup        bool      `json:"isGroup"`
	Name           *string   `json:"name"`
	CreatedBy      *string   `json:"createdBy"`
	CreatedAt      time.Time `json:"createdAt"`
}

func NewConversationResponse(c *models.Conversation) *ConversationResponse {
	return &ConversationResponse{
		ConversationID: c.ID,
		IsGroup:        c.IsGroup,
		Name:           c.Name,
		CreatedBy:      c.CreatedBy,
		CreatedAt:      c.CreatedAt,
	}
}
