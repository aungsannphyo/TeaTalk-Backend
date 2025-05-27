package response

import (
	"time"
)

type MessageResponse struct {
	MessageID        string    `json:"messageId"`
	TargetID         string    `json:"targetId"`
	SenderID         string    `json:"senderId"`
	Content          string    `json:"content"`
	IsRead           bool      `json:"isRead"`
	SeenByName       *string   `json:"seenByName,omitempty"`
	MessageCreatedAt time.Time `json:"messageCreatedAt"`
}
