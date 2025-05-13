package models

import "time"

type ConversationMember struct {
	ConversationID string    `json:"conversation_id"`
	UserID         string    `json:"user_id"`
	JoinedAt       time.Time `json:"joined_at"`
}
