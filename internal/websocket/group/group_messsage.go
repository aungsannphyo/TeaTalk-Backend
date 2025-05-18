package group

import "time"

type WSGroupMessage struct {
	GroupID   string    `json:"group_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
