package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type ConversationService interface {
	CreateGroup(c *gin.Context, dto dto.CreateGroupDto) error
	UpdateGroupName(c *gin.Context, dto dto.UpdateGroupNameDto) error
	InviteGroup(c *gin.Context, dto dto.InviteGroupDto) error
	ModerateGroupInvite(c *gin.Context, dto dto.ModerateGroupInviteDto) error
	AssignAdmin(c *gin.Context, dto dto.AssignAdminDto) error
	GetGroupMembers(ctx context.Context, conversationId string) ([]models.User, error)
}
