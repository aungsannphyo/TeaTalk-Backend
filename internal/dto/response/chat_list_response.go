package response

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type ChatListResponse struct {
	ConversationID       string  `json:"conversationId"`
	IsGroup              bool    `json:"isGroup"`
	Name                 string  `json:"name"`
	LastMessageID        *string `json:"lastMessageId"`
	LastMessageContent   *string `json:"lastMessageContent"`
	LastMessageSender    *string `json:"lastMessageSender"`
	LastMessageCreatedAt *string `json:"lastMessageCreatedAt"`
	UnreadCount          int     `json:"unreadCount"`
	ReceiverID           *string `json:"receiverId"`
	ProfileImage         *string `json:"image"`
	TotalOnline          int     `json:"totalOnline"`
	LastSeen             *string `json:"lastSeen"`
}

func NewChatListResponse(cl *models.ChatListItem) *ChatListResponse {
	return &ChatListResponse{
		ConversationID:       cl.ConversationID,
		IsGroup:              cl.IsGroup,
		Name:                 cl.Name,
		LastMessageID:        cl.LastMessageID,
		LastMessageContent:   cl.LastMessageContent,
		LastMessageSender:    cl.LastMessageSender,
		LastMessageCreatedAt: cl.LastMessageCreatedAt,
		UnreadCount:          cl.UnreadCount,
		ReceiverID:           cl.ReceiverID,
		ProfileImage:         cl.ProfileImage,
		TotalOnline:          cl.TotalOnline,
		LastSeen:             cl.LastSeen,
	}
}
