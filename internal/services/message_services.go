package services

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/google/uuid"
)

type messageService struct {
	mRepo  r.MessageRepository
	fRepo  r.FriendRepository
	cRepo  r.ConversationRepository
	cmRepo r.ConversationMemeberRepository
}

func (s *messageService) SendPrivateMessage(
	ctx context.Context,
	senderID string,
	dto dto.SendPrivateMessageDto,
) error {

	// Step 1: Ensure sender and receiver are friends
	if !s.fRepo.AlreadyFriends(ctx, senderID, dto.ReceiverId) {
		return &e.UnAuthorizedError{Message: "You can't send the message right now!"}
	}

	// Step 2: Check for existing private conversation
	conversation, err := s.cRepo.CheckExistsConversation(ctx, senderID, dto.ReceiverId)
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

		// Add both users as members
		senderMember := &models.ConversationMember{
			ConversationID: conversationID,
			UserID:         senderID,
		}
		receiverMember := &models.ConversationMember{
			ConversationID: conversationID,
			UserID:         dto.ReceiverId,
		}

		if err := s.cmRepo.CreateConversationMember(senderMember); err != nil {
			return &e.InternalServerError{Message: "Failed to add sender to conversation"}
		}
		if err := s.cmRepo.CreateConversationMember(receiverMember); err != nil {
			return &e.InternalServerError{Message: "Failed to add receiver to conversation"}
		}
	} else {
		conversationID = conversation.ID
	}

	// Step 3: Create and save the message
	message := &models.Message{
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        dto.Content,
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

	message := &models.Message{
		ConversationID: cID,
		SenderID:       userID,
		Content:        dto.Content,
	}

	return s.mRepo.CreateMessage(message)
}
