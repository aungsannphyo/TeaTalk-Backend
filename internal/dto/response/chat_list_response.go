package response

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
