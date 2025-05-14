package service

import (
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gin-gonic/gin"
)

type ConversationService interface {
	CreateGroup(dto dto.CreateGroupDto, c *gin.Context) error
	UpdateGroupName(dto dto.UpdateGroupNameDto, c *gin.Context) error
	InviteGroup(dto dto.InviteGroupDto, c *gin.Context) error
	ModerateGroupInvite(dto dto.ModerateGroupInviteDto, c *gin.Context) error
	AssignAdmin(dto dto.AssignAdminDto, c *gin.Context) error
}
