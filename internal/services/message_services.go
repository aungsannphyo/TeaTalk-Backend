package services

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type messageService struct {
	mRepo  r.MessageRepository
	fRepo  r.FriendRepository
	cRepo  r.ConversationRepository
	cmRepo r.ConversationMemeberRepository
}

func (s *messageService) SendPrivateMessage(c *gin.Context, dto dto.SendPrivateMessageDto) error {
	senderId := c.GetString("userId")

	// Step 1: Ensure sender and receiver are friends
	if !s.fRepo.AlreadyFriends(c.Request.Context(), senderId, dto.ReceiverId) {
		return &common.UnAuthorizedError{Message: "You can't send the message right now!"}
	}

	// Step 2: Check for existing private conversation
	conversations, err := s.cRepo.CheckExistsConversation(c.Request.Context(), senderId, dto.ReceiverId)
	if err != nil {
		return &common.InternalServerError{Message: err.Error()}
	}

	var conversationID string
	if len(conversations) != 2 {
		// Create new conversation
		conversationID = uuid.NewString()
		conversation := &models.Conversation{
			ID:        conversationID,
			IsGroup:   false,
			Name:      nil,
			CreatedBy: nil,
		}

		if err := s.cRepo.CreateConversation(conversation); err != nil {
			return &common.InternalServerError{Message: err.Error()}
		}

		// Add both users as members
		senderMember := &models.ConversationMember{
			ConversationID: conversationID,
			UserID:         senderId,
		}
		receiverMember := &models.ConversationMember{
			ConversationID: conversationID,
			UserID:         dto.ReceiverId,
		}

		if err := s.cmRepo.CreateConversationMember(senderMember); err != nil {
			return &common.InternalServerError{Message: "Failed to add sender to conversation"}
		}
		if err := s.cmRepo.CreateConversationMember(receiverMember); err != nil {
			return &common.InternalServerError{Message: "Failed to add receiver to conversation"}
		}
	} else {
		// Reuse existing conversation ID
		conversationID = conversations[0].ID
	}

	// Step 3: Create and save the message
	message := &models.Message{
		ConversationID: conversationID,
		SenderID:       senderId,
		Content:        dto.Content,
	}

	return s.mRepo.CreateMessage(message)
}

func (s *messageService) SendGroupMessage(c *gin.Context, dto dto.SendGroupMessageDto) error {
	cID := c.Param("groupId")
	userId := c.GetString("userId")

	var conversation models.Conversation
	conversation.ID = cID

	con := s.cRepo.CheckExistsGroup(c.Request.Context(), &conversation)

	var cMember models.ConversationMember
	cMember.ConversationID = cID
	cMember.UserID = userId

	member := s.cmRepo.CheckConversationMember(c.Request.Context(), &cMember)

	if !con || !member {
		return &common.ForbiddenError{Message: "You are not a member of this group."}
	}

	message := &models.Message{
		ConversationID: cID,
		SenderID:       userId,
		Content:        dto.Content,
	}

	return s.mRepo.CreateMessage(message)
}
