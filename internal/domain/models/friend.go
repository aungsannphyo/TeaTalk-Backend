package models

import "time"

type Friend struct {
	UserID    string    `json:"user_id"`
	FriendID  string    `json:"friend_id"`
	CreatedAt time.Time `json:"created_at"`
}
