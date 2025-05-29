package models

type ConversationKey struct {
	ConversationId       string `json:"conversation_id"`
	UserID               string `json:"user_id"`
	EncryptedKey         []byte `json:"encrypted_key"`
	ConversationKeyNonce []byte `json:"nonce"`
}
