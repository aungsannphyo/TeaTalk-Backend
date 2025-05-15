package websocket

type WSMessage struct {
	Type      string `json:"type"`                 // "private", "group", "join_group", "leave_group"
	ToUserID  string `json:"to_user_id,omitempty"` // for private messages
	GroupID   string `json:"group_id,omitempty"`   // for group messages
	Content   string `json:"content,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
