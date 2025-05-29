package services

import (
	"context"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/google/uuid"
)

type messageService struct {
	mRepo      r.MessageRepository
	fRepo      r.FriendRepository
	cRepo      r.ConversationRepository
	cmRepo     r.ConversationMemeberRepository
	convKeySvc s.ConversationKeyService
	cKeyRepo   r.ConversationKeyRepository
}

func (s *messageService) SendPrivateMessage(
	ctx context.Context,
	senderID string,
	dto dto.SendPrivateMessageDto,
) error {
	// Step 1: Ensure sender and receiver are friends
	if !s.fRepo.AlreadyFriends(ctx, senderID, dto.ReceiverID) {
		return &e.UnAuthorizedError{Message: "You can't send the message right now!"}
	}

	// Step 2: Check for existing private conversation
	conversation, err := s.cRepo.CheckExistsConversation(ctx, senderID, dto.ReceiverID)

	if err != nil {
		return &e.InternalServerError{Message: err.Error()}
	}

	var conversationID string
	if conversation.ID == "" {
		// Create new conversation
		conversationID = uuid.NewString()
		conversation := &models.Conversation{
			ID:        conversationID,
			IsGroup:   false,
			Name:      nil,
			CreatedBy: nil,
		}

		if err := s.cRepo.CreateConversation(conversation); err != nil {

			return &e.InternalServerError{Message: err.Error()}
		}

		// store conversation key for both sender and receiver
		userIDs := []string{senderID, dto.ReceiverID}
		if err := s.convKeySvc.EncryptConversationForUsers(ctx, userIDs, conversationID); err != nil {
			return err
		}

		// Add both users as members
		members := []*models.ConversationMember{
			{
				ConversationID: conversationID,
				UserID:         senderID,
			},
			{
				ConversationID: conversationID,
				UserID:         dto.ReceiverID,
			},
		}

		//create conversation for both sender and receiver
		for _, member := range members {
			if err := s.cmRepo.CreateConversationMember(member); err != nil {
				return &e.InternalServerError{Message: "Failed to add user to conversation"}
			}
		}
	} else {
		conversationID = conversation.ID
	}

	//  Load user's encrypted conversation key + nonce from DB
	encryptedKey, nonce, err := s.cKeyRepo.GetConversationKey(ctx, conversationID, senderID)
	if err != nil {
		return &e.InternalServerError{Message: "Failed to add retrieve to conversation key"}
	}

	encryptedMessage, messageNonce, err := s.convKeySvc.EncryptMessage(ctx, encryptedKey, nonce, senderID, dto.Content)

	if err != nil {
		return err
	}

	messageID := uuid.New().String()
	// Step 3: Create and save the message
	message := &models.Message{
		ID:             messageID,
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        encryptedMessage,
		MessageNonce:   messageNonce,
	}

	return s.mRepo.CreateMessage(message)
}

func (s *messageService) SendGroupMessage(
	ctx context.Context,
	cID string,
	userID string,
	dto dto.SendGroupMessageDto,
) error {

	var conversation models.Conversation
	conversation.ID = cID

	con := s.cRepo.CheckExistsGroup(ctx, &conversation)

	var cMember models.ConversationMember
	cMember.ConversationID = cID
	cMember.UserID = userID

	member := s.cmRepo.CheckConversationMember(ctx, &cMember)

	if !con || !member {
		return &e.ForbiddenError{Message: "You are not a member of this group."}
	}

	messageID := uuid.New().String()

	message := &models.Message{
		ID:             messageID,
		ConversationID: cID,
		SenderID:       userID,
		Content:        []byte(dto.Content),
	}

	return s.mRepo.CreateMessage(message)
}

func (s *messageService) GetMessages(
	ctx context.Context,
	conversationID string,
	userID string,
	cursorTimestamp *time.Time,
	pageSize int,
) ([]response.MessageResponse, error) {
	messages, err := s.mRepo.GetMessages(ctx, conversationID, cursorTimestamp, pageSize)

	if err != nil {
		return nil, &e.InternalServerError{Message: "Something went wrong, please try again later"}
	}

	msgMap := map[string]response.MessageResponse{}
	for _, msg := range messages {
		// Skip logged-in user's perspective row
		if msg.MemberID == userID {
			continue
		}

		plainContentBytes, err := s.convKeySvc.DecryptMessage(ctx, userID, conversationID, []byte(msg.Content), msg.MessageNonce)

		if err != nil {
			return nil, &e.InternalServerError{Message: "Something went wrong, please try again later"}
		}

		msgMap[msg.MessageID] = response.MessageResponse{
			MessageID:        msg.MessageID,
			TargetID:         msg.MemberID,
			SenderID:         msg.SenderID,
			Content:          string(plainContentBytes),
			IsRead:           msg.IsRead,
			SeenByName:       msg.SeenByName,
			MessageCreatedAt: msg.MessageCreatedAt,
		}
	}

	msgResponse := make([]response.MessageResponse, 0, len(msgMap))
	for _, m := range msgMap {
		msgResponse = append(msgResponse, m)
	}

	return msgResponse, nil
}
