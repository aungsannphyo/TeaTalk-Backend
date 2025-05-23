package models

type ChatListItem struct {
	ConversationID       string `json:"conversation_id"`
	IsGroup              bool   `json:"is_group"`
	Name                 string `json:"name"`
	LastMessageID        string `json:"last_message_id"`
	LastMessageContent   string `json:"last_message_content"`
	LastMessageSender    string `json:"last_message_sender"`
	LastMessageCreatedAt string `json:"last_message_created_at"`
	UnreadCount          int    `json:"unread_count"`
}
