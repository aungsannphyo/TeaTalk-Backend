package models

import "time"

type FriendRequestAction string

const (
	ActionSend       FriendRequestAction = "SENT"
	ActionAccepted   FriendRequestAction = "ACCEPTED"
	ActionRejected   FriendRequestAction = "REJECTED"
	ActionCancelled  FriendRequestAction = "CANCELLED"
	ActionUnFriended FriendRequestAction = "UNFRIENDED"
)

type FriendRequestLog struct {
	ID          string              `json:"id"`
	SenderID    string              `json:"sender_id"`
	ReceiverID  string              `json:"receiver_id"`
	Action      FriendRequestAction `json:"action"`
	PerformedBy string              `json:"performed_by"`
	CreatedAt   time.Time           `json:"created_at"`
}
