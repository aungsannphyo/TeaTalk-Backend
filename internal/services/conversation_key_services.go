package services

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
)

type cKeyService struct {
	cKeyRepo        r.ConversationKeyRepository
	userRepo        r.UserRepository
	sessionKeyCache *SessionKeyCache
}

func (s *cKeyService) EncryptConversationForUsers(
	ctx context.Context,
	userIDs []string,
	conversationID string,
) error {

	//creation part of conversation key for encrypt
	// 1. Generate conversation symmetric key
	convKey, err := utils.GenerateRandomBytes(utils.KeyLength)
	if err != nil {
		return &e.InternalServerError{Message: "Failed to generate random bytes"}
	}

	// 2. For each user, encrypt the convKey with user's decrypted user key
	for _, userID := range userIDs {
		decryptUserKey, found := s.sessionKeyCache.GetUserDecryptedKey(userID)

		if !found {
			return &e.InternalServerError{Message: "user key not in session — must reauthenticate"}
		}

		// 3. Store encryptedKey + nonce for (conversationID, userID) in conversation_keys table
		encryptedKey, nonce, err := utils.EncryptConversationKey(convKey, decryptUserKey)

		if err != nil {
			return &e.InternalServerError{Message: "Failed to encrypt by sender key"}
		}

		ckey := &models.ConversationKey{
			UserID:               userID,
			ConversationId:       conversationID,
			EncryptedKey:         encryptedKey,
			ConversationKeyNonce: nonce,
		}
		if err := s.cKeyRepo.CreateConversationKey(ckey); err != nil {
			return &e.InternalServerError{Message: "Failed to add  key to conversation key "}
		}
	}

	return nil
}

func (s *cKeyService) EncryptMessage(ctx context.Context,
	encryptedKey []byte,
	nonce []byte,
	content string,
	userID string,
) ([]byte, []byte, error) {

	// 1. Decrypt conversation key with user's decrypted user key
	userKey, found := s.sessionKeyCache.GetUserDecryptedKey(userID)

	if !found {
		return nil, nil, &e.InternalServerError{Message: "user key not in session — must reauthenticate"}
	}
	convKey, err := utils.DecryptConversationKey(encryptedKey, nonce, userKey)
	if err != nil {
		return nil, nil, &e.InternalServerError{Message: "Failed to get by decrypt conversation key"}
	}

	// 2. Encrypt message content with conversation key
	plainText := []byte(content)
	encryptedMessage, messageNonce, err := utils.EncryptMessageContent(plainText, convKey)
	if err != nil {
		return nil, nil, err
	}

	return encryptedMessage, messageNonce, nil
}

func (s *cKeyService) DecryptMessage(
	ctx context.Context,
	userID,
	conversationID string,
	encryptedMessage,
	messageNonce []byte,
) ([]byte, error) {
	encryptKey, nonce, err := s.cKeyRepo.GetConversationKey(ctx, conversationID, userID)

	if err != nil {
		return nil, &e.InternalServerError{Message: "Failed to get by decrypt conversation key"}
	}

	userKey, found := s.sessionKeyCache.GetUserDecryptedKey(userID)

	if !found {
		return nil, &e.InternalServerError{Message: "user key not in session — must reauthenticate"}
	}

	convKey, err := utils.DecryptConversationKey(encryptKey, nonce, userKey)
	if err != nil {
		return nil, err
	}

	return utils.DecryptMessageContent(encryptedMessage, nonce, convKey)
}
