package services

import (
	"context"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/google/uuid"
)

type messageService struct {
	mRepo  r.MessageRepository
	fRepo  r.FriendRepository
	cRepo  r.ConversationRepository
	cmRepo r.ConversationMemeberRepository
	mrRepo r.MessageReadRepository
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

	messageID := uuid.New().String()
	// Step 3: Create and save the message
	message := &models.Message{
		ID:             messageID,
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        dto.Content,
	}

	//Step 4: Create Message Read Default for Sender and save the message reads
	msgRead := &models.MessageRead{
		MessageId: messageID,
		UserID:    senderID,
	}
	err = s.mrRepo.CreateReadMessage(msgRead)
	if err != nil {
		return &e.InternalServerError{Message: "Failed to add sender to message read table"}
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
		Content:        dto.Content,
	}

	msgRead := &models.MessageRead{
		MessageId: messageID,
		UserID:    userID,
	}
	err := s.mrRepo.CreateReadMessage(msgRead)
	if err != nil {
		return &e.InternalServerError{Message: "Failed to add sender to message read table"}
	}

	return s.mRepo.CreateMessage(message)
}

func (s *messageService) GetMessages(
	ctx context.Context,
	conversationID string,
	cursorTimestamp *time.Time,
	pageSize int,
) ([]response.MessageResponse, error) {
	messages, err := s.mRepo.GetMessages(ctx, conversationID, cursorTimestamp, pageSize)

	if err != nil {
		return nil, &e.InternalServerError{Message: "Something went wrong, please try again later"}
	}

	return messages, nil

}
