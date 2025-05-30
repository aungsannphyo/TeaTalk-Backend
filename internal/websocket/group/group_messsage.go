package group

import "time"

type WSGroupMessage struct {
	ConversationID string    `json:"conversationId,omitempty"`
	SenderID       string    `json:"senderID,omitempty"`
	Content        string    `json:"content,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
