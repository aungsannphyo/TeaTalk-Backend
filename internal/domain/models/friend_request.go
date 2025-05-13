package models

import "time"

type FriendRequestStatus string

const (
	StatusPending  FriendRequestStatus = "PENDING"
	StatusAccepted FriendRequestStatus = "ACCEPTED"
	StatusRejected FriendRequestStatus = "REJECTED"
)

type FriendRequest struct {
	ID         string              `json:"id"`
	SenderId   string              `json:"sender_id"`
	ReceiverId string              `json:"receiver_id"`
	Status     FriendRequestStatus `json:"status"`
	CreatedAt  time.Time           `json:"created_at"`
}
