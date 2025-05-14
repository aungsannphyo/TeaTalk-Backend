package models

import "time"

type GroupAdmin struct {
	ConversationID string    `json:"conversation_id"`
	UserID         string    `json:"user_id"`
	AssignedAt     time.Time `json:"assigned_at"`
}
