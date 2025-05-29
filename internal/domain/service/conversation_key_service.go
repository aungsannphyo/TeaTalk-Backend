package service

import "context"

type ConversationKeyService interface {
	EncryptConversationForUsers(
		ctx context.Context,
		userIDs []string,
		conversationID string,
	) error
	EncryptMessage(
		ctx context.Context,
		encryptedKey []byte,
		nonce []byte,
		content string,
		userID string,
	) ([]byte, []byte, error)
	DecryptMessage(
		ctx context.Context,
		userID,
		conversationID string,
		encryptedMessage,
		messageNonce []byte,
	) ([]byte, error)
}
