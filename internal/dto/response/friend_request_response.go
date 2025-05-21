package response

import "time"

type FriendRequestResponse struct {
	RequestID    string    `json:"requestId"`
	SenderID     string    `json:"senderId"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	ProfileImage string    `json:"profileImage"`
	CreateAt     time.Time `json:"createAt"`
}
