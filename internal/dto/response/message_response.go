package response

import (
	"time"
)

type MessageResponse struct {
	ID               string    `json:"id"`
	ConversationID   string    `json:"conversationId"`
	SenderID         string    `json:"senderId"`
	ReceiverID       string    `json:"receiverId"`
	Content          string    `json:"content"`
	IsRead           bool      `json:"isRead"`
	SeenByName       *string   `json:"seenByName,omitempty"`
	MessageCreatedAt time.Time `json:"messageCreatedAt"`
}
