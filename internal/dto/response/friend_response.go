package response

import (
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type FriendResponse struct {
	ID             string     `json:"id"`
	UserIdentity   string     `json:"userIdentity"`
	Username       string     `json:"username"`
	Email          string     `json:"email"`
	ProfileImage   *string    `json:"profileImage"`
	IsOnline       *bool      `json:"isOnline"`
	LastSeen       *time.Time `json:"lastSeen"`
	ConversationId *string    `json:"conversationId,omitempty"`
}

func NewFriendResponse(user *models.User, ps *models.PersonalDetails) *FriendResponse {
	return &FriendResponse{
		ID:           user.ID,
		UserIdentity: user.UserIdentity,
		Username:     user.Username,
		Email:        user.Email,
		ProfileImage: ps.ProfileImage,
		IsOnline:     &ps.IsOnline,
		LastSeen:     &ps.LastSeen,
	}
}
