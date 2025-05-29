package models

import "time"

type Message struct {
	ID             string    `json:"id"`
	ConversationID string    `json:"conversation_id"`
	SenderID       string    `json:"sender_id"`
	Content        []byte    `json:"content"`
	MessageNonce   []byte    `json:"message_nonce"`
	CreatedAt      time.Time `json:"created_at"`
}
