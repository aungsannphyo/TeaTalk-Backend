package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type ConversationService interface {
	CreateConversation(userID string, dto dto.CreateConversationDto) error
	UpdateGroupName(conversationID string, dto dto.UpdateGroupNameDto) error
	InviteGroup(ctx context.Context, conversationID string, userID string, dto dto.InviteGroupDto) error
	ModerateGroupInvite(ctx context.Context, conversationID string, inviteID string, userID string, dto dto.ModerateGroupInviteDto) error
	AssignAdmin(ctx context.Context, conversationID string, userID string, dto dto.AssignAdminDto) error
	GetGroupMembers(ctx context.Context, conversationId string) ([]models.User, error)
	GetGroupsById(ctx context.Context, userID string) ([]models.Conversation, error)
}
