package handler

import (
	"fmt"

	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
)

type ConversationsHandler struct {
	cService s.ConversationService
}

func NewConversationHandler(s s.ConversationService) *ConversationsHandler {
	return &ConversationsHandler{
		cService: s,
	}
}

func (s *ConversationsHandler) CreateGroup(c *gin.Context) {
	var cgDto dto.CreateGroupDto

	if err := c.ShouldBindJSON(&cgDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateCreateGroup(cgDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.CreateGroup(c, cgDto); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.CreateResponse(c, "You have been successfully created the group!")
}

func (s *ConversationsHandler) UpdateGroupName(c *gin.Context) {
	var ugDto dto.UpdateGroupNameDto

	if err := c.ShouldBindJSON(&ugDto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.UpdateGroupName(c, ugDto); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{"message": "You have been successfully updated the group name!"})
}

func (s *ConversationsHandler) InviteGroup(c *gin.Context) {
	var igdto dto.InviteGroupDto

	if err := c.ShouldBindJSON(&igdto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateInviteGroup(igdto); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.InviteGroup(c, igdto); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{"message": "You have been successfully invited that you selected users!"})
}

func (s *ConversationsHandler) ModerateGroupInvite(c *gin.Context) {
	var mgi dto.ModerateGroupInviteDto

	if err := c.ShouldBindJSON(&mgi); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateModerateGroupInvite(mgi); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.ModerateGroupInvite(c, mgi); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{
		"message": fmt.Sprintf("You have been successfully %v!", mgi.Status),
	})
}

func (s *ConversationsHandler) AssignAdmin(c *gin.Context) {
	var aa dto.AssignAdminDto

	if err := c.ShouldBindJSON(&aa); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateAssignAdmin(aa); err != nil {
		common.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.AssignAdmin(c, aa); err != nil {
		common.InternalServerResponse(c, err)
		return
	}

	common.OkResponse(c, gin.H{"message": "You have been successfully made this action!"})
}
