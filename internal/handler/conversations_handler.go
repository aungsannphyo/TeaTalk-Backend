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
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&cgDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateCreateGroup(cgDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.CreateGroup(userID, cgDto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.CreateResponse(c, "You have been successfully created the group!")
}

func (s *ConversationsHandler) UpdateGroupNameHandler(c *gin.Context) {
	var ugDto dto.UpdateGroupNameDto

	groupID := c.Param("groupID")

	if err := c.ShouldBindJSON(&ugDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.UpdateGroupName(groupID, ugDto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have been successfully updated the group name!"})
}

func (s *ConversationsHandler) InviteGroupHandler(c *gin.Context) {
	var igdto dto.InviteGroupDto
	groupID := c.Param("groupID")
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&igdto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateInviteGroup(igdto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.InviteGroup(c.Request.Context(), groupID, userID, igdto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have been successfully invited that you selected users!"})
}

func (s *ConversationsHandler) ModerateGroupInviteHandler(c *gin.Context) {
	var mgi dto.ModerateGroupInviteDto

	userID := c.GetString("userID")
	groupID := c.Param("groupID")
	inviteUserID := c.Param("inviteUserID")

	if err := c.ShouldBindJSON(&mgi); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateModerateGroupInvite(mgi); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.ModerateGroupInvite(c.Request.Context(), groupID, inviteUserID, userID, mgi); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{
		"message": fmt.Sprintf("You have been successfully %v!", mgi.Status),
	})
}

func (s *ConversationsHandler) AssignAdminHandler(c *gin.Context) {
	var aa dto.AssignAdminDto
	userID := c.GetString("userID")
	groupID := c.Param("groupID")

	if err := c.ShouldBindJSON(&aa); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateAssignAdmin(aa); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := s.cService.AssignAdmin(c.Request.Context(), groupID, userID, aa); err != nil {
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

func (h *ConversationsHandler) GetGroupsByIdHandler(c *gin.Context) {
	userID := c.GetString("userID")
	groups, err := h.cService.GetGroupsById(c.Request.Context(), userID)

	if err != nil {
		e.NotFoundResponse(c, err)
		return
	}

	var groupList []response.Conversation

	if len(groupList) == 0 {
		success.OkResponse(c, []models.Conversation{})
	} else {
		for _, c := range groups {
			conversation := &models.Conversation{
				ID:        c.ID,
				IsGroup:   c.IsGroup,
				Name:      c.Name,
				CreatedBy: c.CreatedBy,
				CreatedAt: c.CreatedAt,
			}
			cResponse := response.NewConversationResponse(conversation)
			groupList = append(groupList, *cResponse)
		}

		success.OkResponse(c, groupList)
	}
}
