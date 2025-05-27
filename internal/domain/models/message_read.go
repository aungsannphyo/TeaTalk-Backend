package models

import "time"

type MessageRead struct {
	MessageId string    `json:"message_id"`
	UserID    string    `json:"user_id"`
	ReadAt    time.Time `json:"read_at"`
}
