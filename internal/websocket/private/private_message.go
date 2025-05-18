package private

import "time"

type WSPrivateMessage struct {
	ReceiverID string    `json:"receiver_id,omitempty"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
