package dto

import (
	"strings"

	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type CreateConversationKeyDto struct {
	ConversationId           string `json:"conversationId"`
	UserID                   string `json:"userId"`
	ConversationEncryptedKey string `json:"conversationEncryptedKey"`
	ConversationKeyNonce     string `json:"conversationKeyNonce"`
}

func ValidateCreateConversationKey(ck CreateConversationKeyDto) error {
	var errs v.ValidationErrors

	if strings.TrimSpace(ck.ConversationId) == "" {
		errs = append(errs, v.ValidationError{Field: "Conversation Id", Message: "Conversation Id is required"})
	}

	if strings.TrimSpace(ck.UserID) == "" {
		errs = append(errs, v.ValidationError{Field: "User Id", Message: "User Id is required"})
	}

	if strings.TrimSpace(ck.ConversationEncryptedKey) == "" {
		errs = append(errs, v.ValidationError{Field: "Conversation Encrypted Key", Message: "Conversation Encrypted Key is required"})
	}

	if strings.TrimSpace(ck.ConversationKeyNonce) == "" {
		errs = append(errs, v.ValidationError{Field: "Conversation key nonce", Message: "Conversation key nonce is required"})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
