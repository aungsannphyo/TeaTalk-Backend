package private

import "time"

type WSPrivateMessage struct {
	ReceiverID string    `json:"receiverId,omitempty"`
	SenderID   string    `json:"senderID,omitempty"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
