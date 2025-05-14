package models

import "time"

type GroupInviteStatus string

const (
	GroupPending  GroupInviteStatus = "PENDING"
	GroupRejected GroupInviteStatus = "REJECTED"
	GroupApproved GroupInviteStatus = "APPROVED"
)

type GroupInvite struct {
	ID             string            `json:"id"`
	ConversationID string            `json:"conversation_id"`
	InvitedBy      string            `json:"invited_by"`
	InvitedUserId  string            `json:"invited_user_id"`
	Status         GroupInviteStatus `json:"status"`
	CreatedAt      time.Time         `json:"created_at"`
}
