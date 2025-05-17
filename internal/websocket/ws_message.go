package websocket

import "time"

type WSMessage struct {
	Type      string    `json:"type"`                 // "private", "group",
	ToUserID  string    `json:"to_user_id,omitempty"` // for private messages
	GroupID   string    `json:"group_id,omitempty"`   // for group messages
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
