package models

import "time"

type FriendRequestStatus string

const (
	FriendRequestPending  FriendRequestStatus = "PENDING"
	FriendRequestAccepted FriendRequestStatus = "ACCEPTED"
	FriendRequestRejected FriendRequestStatus = "REJECTED"
)

type FriendRequest struct {
	ID         string              `json:"id"`
	SenderId   string              `json:"sender_id"`
	ReceiverId string              `json:"receiver_id"`
	Status     FriendRequestStatus `json:"status"`
	CreatedAt  time.Time           `json:"created_at"`
}
