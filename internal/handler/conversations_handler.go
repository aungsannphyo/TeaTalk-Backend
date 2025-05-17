package handler

import (
	"fmt"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/aungsannphyo/ywartalk/pkg/success"
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

func (s *ConversationsHandler) CreateGroupHandler(c *gin.Context) {
	var cgDto dto.CreateGroupDto

	if err := c.ShouldBindJSON(&cgDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateCreateGroup(cgDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.CreateGroup(c, cgDto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.CreateResponse(c, "You have been successfully created the group!")
}

func (s *ConversationsHandler) UpdateGroupNameHandler(c *gin.Context) {
	var ugDto dto.UpdateGroupNameDto

	if err := c.ShouldBindJSON(&ugDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.UpdateGroupName(c, ugDto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have been successfully updated the group name!"})
}

func (s *ConversationsHandler) InviteGroupHandler(c *gin.Context) {
	var igdto dto.InviteGroupDto

	if err := c.ShouldBindJSON(&igdto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateInviteGroup(igdto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.InviteGroup(c, igdto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have been successfully invited that you selected users!"})
}

func (s *ConversationsHandler) ModerateGroupInviteHandler(c *gin.Context) {
	var mgi dto.ModerateGroupInviteDto

	if err := c.ShouldBindJSON(&mgi); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateModerateGroupInvite(mgi); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.ModerateGroupInvite(c, mgi); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{
		"message": fmt.Sprintf("You have been successfully %v!", mgi.Status),
	})
}

func (s *ConversationsHandler) AssignAdminHandler(c *gin.Context) {
	var aa dto.AssignAdminDto

	if err := c.ShouldBindJSON(&aa); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateAssignAdmin(aa); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.AssignAdmin(c, aa); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have been successfully made this action!"})
}

func (h *ConversationsHandler) GetGroupMembersHandler(c *gin.Context) {
	groupId := c.Param("groupId")

	groupUsers, err := h.cService.GetGroupMembers(c.Request.Context(), groupId)

	if err != nil {
		e.InternalServerResponse(c, err)
	}

	var users []response.UserResponse

	for _, groupUser := range groupUsers {
		user := &models.User{
			ID:        groupUser.ID,
			Email:     groupUser.Email,
			Username:  groupUser.Username,
			CreatedAt: groupUser.CreatedAt,
		}
		userResponse := response.NewUserResponse(user)
		users = append(users, *userResponse)
	}

	success.OkResponse(c, users)
}
