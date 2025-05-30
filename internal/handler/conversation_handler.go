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

type ConversationHandler struct {
	cService s.ConversationService
}

func NewConversationHandler(s s.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		cService: s,
	}
}

func (h *ConversationHandler) UpdateGroupNameHandler(c *gin.Context) {
	var ugDto dto.UpdateGroupNameDto

	conversationID := c.Param("conversationID")

	if err := c.ShouldBindJSON(&ugDto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.cService.UpdateGroupName(conversationID, ugDto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have been successfully updated the group name!"})
}

func (h *ConversationHandler) InviteGroupHandler(c *gin.Context) {
	var igdto dto.InviteGroupDto
	conversationID := c.Param("conversationID")
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&igdto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateInviteGroup(igdto); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.cService.InviteGroup(c.Request.Context(), conversationID, userID, igdto); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have been successfully invited that you selected users!"})
}

func (h *ConversationHandler) ModerateGroupInviteHandler(c *gin.Context) {
	var mgi dto.ModerateGroupInviteDto

	userID := c.GetString("userID")
	conversationID := c.Param("conversationID")
	inviteUserID := c.Param("inviteUserID")

	if err := c.ShouldBindJSON(&mgi); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateModerateGroupInvite(mgi); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.cService.ModerateGroupInvite(c.Request.Context(), conversationID, inviteUserID, userID, mgi); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{
		"message": fmt.Sprintf("You have been successfully %v!", mgi.Status),
	})
}

func (h *ConversationHandler) AssignAdminHandler(c *gin.Context) {
	var aa dto.AssignAdminDto
	userID := c.GetString("userID")
	conversationID := c.Param("conversationID")

	if err := c.ShouldBindJSON(&aa); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := dto.ValidateAssignAdmin(aa); err != nil {
		e.BadRequestResponse(c, err)
		return
	}

	if err := h.cService.AssignAdmin(c.Request.Context(), conversationID, userID, aa); err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	success.OkResponse(c, gin.H{"message": "You have been successfully made this action!"})
}

func (h *ConversationHandler) GetGroupMembersHandler(c *gin.Context) {
	conversationID := c.Param("conversationID")

	groupUsers, err := h.cService.GetGroupMembers(c.Request.Context(), conversationID)

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

func (h *ConversationHandler) GetGroupsByIdHandler(c *gin.Context) {
	userID := c.GetString("userID")
	groups, err := h.cService.GetGroupsById(c.Request.Context(), userID)

	if err != nil {
		e.NotFoundResponse(c, err)
		return
	}

	if len(groups) == 0 {
		success.OkResponse(c, []models.Conversation{})
	} else {

		success.OkResponse(c, groups)
	}
}

func (h *ConversationHandler) GetConversationHandler(c *gin.Context) {
	senderID := c.Query("senderId")
	receiverID := c.Query("receiverId")

	conversation, err := h.cService.GetConversation(c.Request.Context(), senderID, receiverID)

	if err != nil {
		e.InternalServerResponse(c, err)
		return
	}

	cResponse := response.NewConversationResponse(conversation)

	success.OkResponse(c, cResponse)
}
