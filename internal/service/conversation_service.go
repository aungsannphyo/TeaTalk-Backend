package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConversationService struct {
	cRepo  r.ConversationRepository
	cmRepo r.ConversationMemeberRepository
	gaRepo r.GroupAdminRepository
	giRepo r.GroupInviteRepository
	fRepo  r.FriendRepository
}

func NewConversationService(
	cRepo r.ConversationRepository,
	cmRepo r.ConversationMemeberRepository,
	gaRepo r.GroupAdminRepository,
	giRepo r.GroupInviteRepository,
	fRepo r.FriendRepository,
) *ConversationService {
	return &ConversationService{
		cRepo:  cRepo,
		cmRepo: cmRepo,
		gaRepo: gaRepo,
		giRepo: giRepo,
		fRepo:  fRepo,
	}
}

func (s *ConversationService) CreateGroup(dto dto.CreateGroupDto, c *gin.Context) error {
	cID := uuid.NewString()
	userId := c.GetString("userId")

	conversation := &models.Conversation{
		ID:        cID,
		IsGroup:   true,
		Name:      &dto.Name,
		CreatedBy: &userId,
	}

	conversationMember := &models.ConversationMember{
		ConversationID: cID,
		UserID:         userId,
	}

	groupAdmin := &models.GroupAdmin{
		ConversationID: cID,
		UserID:         userId,
	}

	if err := s.cRepo.CreateConversation(conversation); err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if err := s.cmRepo.CreateConversationMember(conversationMember); err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if err := s.gaRepo.CreateGroupAdmin(groupAdmin); err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}
	return nil
}

func (s *ConversationService) UpdateGroupName(dto dto.UpdateGroupNameDto, c *gin.Context) error {
	cID := c.Param("groupId")

	uc := &models.Conversation{
		ID:   cID,
		Name: &dto.Name,
	}

	if err := s.cRepo.UpdateGroupName(uc); err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}
	return nil
}

func (s *ConversationService) InviteGroup(dto dto.InviteGroupDto, c *gin.Context) error {
	//need to check current user is group admin or not
	//check friendship between invitedBy and invited user
	//insert into group_invites with status = "approved" if admin is invite
	//insert into group_invites with status = "pending" if not admin is invite
	//if admin is invite -> Add to conversation_members.

	cID := c.Param("groupId")
	userID := c.GetString("userId")

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(cID, userID)

	if err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	for _, iuser := range dto.InvitedUserId {
		if !s.fRepo.AlreadyFriends(userID, iuser) {
			continue
		}

		status := models.GroupPending
		if isGroupAdmin {
			status = models.GroupApproved
		}

		groupInvite := &models.GroupInvite{
			ConversationID: cID,
			InvitedBy:      userID,
			InvitedUserId:  iuser,
			Status:         status,
		}

		if err := s.giRepo.CreateGroupInvite(groupInvite); err != nil {
			return &common.InternalServerError{Message: "Something went wrong, please try again later"}
		}

		if isGroupAdmin {
			conversationMember := &models.ConversationMember{
				ConversationID: cID,
				UserID:         iuser,
			}
			if err := s.cmRepo.CreateConversationMember(conversationMember); err != nil {
				return &common.InternalServerError{Message: "Something went wrong, please try again later"}
			}
		}

	}

	return nil
}

func (s *ConversationService) ModerateGroupInvite(dto dto.ModerateGroupInviteDto, c *gin.Context) error {
	cID := c.Param("groupId")
	inviteId := c.Param("inviteUserId")
	userID := c.GetString("userId")

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(cID, userID)

	if err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if !isGroupAdmin {
		return &common.ForbiddenError{Message: "You are not an admin of this group"}
	}

	mgi := &models.GroupInvite{
		InvitedUserId:  inviteId,
		ConversationID: cID,
		Status:         dto.Status,
	}

	if err := s.giRepo.ModerateGroupInvite(mgi); err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	conversationMember := &models.ConversationMember{
		ConversationID: cID,
		UserID:         inviteId,
	}

	if err := s.cmRepo.CreateConversationMember(conversationMember); err != nil {
		return &common.InternalServerError{Message: "Something went wrong, please try again later"}
	}

	return nil
}

func (s *ConversationService) AssignAdmin(dto dto.AssignAdminDto, c *gin.Context) error {
	cID := c.Param("groupId")
	userID := c.GetString("userId")

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(cID, userID)

	if err != nil {
		return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if !isGroupAdmin {
		return &common.ForbiddenError{Message: "You are not an admin of this group"}
	}

	for _, iuser := range dto.InvitedUserId {
		groupAdmin := &models.GroupAdmin{
			ConversationID: cID,
			UserID:         iuser,
		}

		if err := s.gaRepo.CreateGroupAdmin(groupAdmin); err != nil {
			return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
		}
	}

	return nil
}
