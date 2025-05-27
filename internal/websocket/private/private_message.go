package private

import "time"

type WSPrivateMessage struct {
	ReceiverID string    `json:"receiverId,omitempty"`
	SenderID   string    `json:"senderId,omitempty"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}
