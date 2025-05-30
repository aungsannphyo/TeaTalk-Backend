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
	mRepo    r.MessageRepository
	fRepo    r.FriendRepository
	cRepo    r.ConversationRepository
	cmRepo   r.ConversationMemeberRepository
	cKeyRepo r.ConversationKeyRepository
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
	} else {
		conversationID = conversation.ID
	}

	messageID := uuid.New().String()
	// Step 3: Create and save the message
	message := &models.Message{
		ID:             messageID,
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        []byte(dto.Content),
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

		msgMap[msg.MessageID] = response.MessageResponse{
			MessageID:        msg.MessageID,
			TargetID:         msg.MemberID,
			SenderID:         msg.SenderID,
			Content:          msg.Content,
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
