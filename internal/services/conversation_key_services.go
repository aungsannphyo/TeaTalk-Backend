package services

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
)

type cKeyService struct {
	cKeyRepo r.ConversationKeyRepository
}

func (s *cKeyService) CreateConversationKey(dto dto.CreateConversationKeyDto) error {

	encryptKey, encryptErr := utils.DecodeBase64(dto.ConversationEncryptedKey)
	nonce, nonceErr := utils.DecodeBase64(dto.ConversationKeyNonce)

	if nonceErr != nil || encryptErr != nil {
		return &e.InternalServerError{Message: "Fail to decode string"}
	}

	cKey := &models.ConversationKey{
		ConversationId:           dto.ConversationId,
		UserID:                   dto.UserID,
		ConversationEncryptedKey: encryptKey,
		ConversationKeyNonce:     nonce,
	}

	if err := s.cKeyRepo.CreateConversationKey(cKey); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	return nil
}
