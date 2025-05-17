package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type ConversationService interface {
	CreateGroup(userID string, dto dto.CreateGroupDto) error
	UpdateGroupName(groupID string, dto dto.UpdateGroupNameDto) error
	InviteGroup(ctx context.Context, groupID string, userID string, dto dto.InviteGroupDto) error
	ModerateGroupInvite(ctx context.Context, groupID string, inviteID string, userID string, dto dto.ModerateGroupInviteDto) error
	AssignAdmin(ctx context.Context, groupID string, userID string, dto dto.AssignAdminDto) error
	GetGroupMembers(ctx context.Context, conversationId string) ([]models.User, error)
}
