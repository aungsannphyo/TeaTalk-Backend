package services

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	e "github.com/aungsannphyo/ywartalk/pkg/error"
	"github.com/google/uuid"
)

type conService struct {
	cRepo  r.ConversationRepository
	cmRepo r.ConversationMemeberRepository
	gaRepo r.GroupAdminRepository
	giRepo r.GroupInviteRepository
	fRepo  r.FriendRepository
}

func (s *conService) addMemberAndAdmin(conversationID, userID string) error {
	member := &models.ConversationMember{
		ConversationID: conversationID,
		UserID:         userID,
	}
	if err := s.cmRepo.CreateConversationMember(member); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	admin := &models.GroupAdmin{
		ConversationID: conversationID,
		UserID:         userID,
	}
	if err := s.gaRepo.CreateGroupAdmin(admin); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	return nil
}

func (s *conService) addConversationMember(conversationID string, memberIDs *[]string) error {
	if memberIDs == nil || len(*memberIDs) == 0 {
		return nil
	}

	for _, memberID := range *memberIDs {
		member := &models.ConversationMember{
			ConversationID: conversationID,
			UserID:         memberID,
		}
		if err := s.cmRepo.CreateConversationMember(member); err != nil {
			return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
		}
	}
	return nil
}

func (s *conService) CreateConversation(userID string, dto dto.CreateConversationDto) error {
	cID := uuid.NewString()

	var c *models.Conversation

	if dto.IsGroup {
		c = &models.Conversation{
			ID:        cID,
			IsGroup:   true,
			Name:      &dto.Name,
			CreatedBy: &userID,
		}
	} else {
		c = &models.Conversation{
			ID:        cID,
			IsGroup:   false,
			Name:      nil,
			CreatedBy: nil,
		}
	}

	// Create the group conversation
	if err := s.cRepo.CreateConversation(c); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	// Add the creator as a member and admin
	//if isGroup is true
	if dto.IsGroup {
		if err := s.addMemberAndAdmin(cID, userID); err != nil {
			return err
		}
	}

	// Add other members (if any)
	if err := s.addConversationMember(cID, dto.MemberID); err != nil {
		return err
	}

	return nil
}

func (s *conService) UpdateGroupName(conversationID string, dto dto.UpdateGroupNameDto) error {

	uc := &models.Conversation{
		ID:   conversationID,
		Name: &dto.Name,
	}

	if err := s.cRepo.UpdateGroupName(uc); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}
	return nil
}

func (s *conService) InviteGroup(ctx context.Context, conversationID string, userID string, dto dto.InviteGroupDto) error {
	//need to check current user is group admin or not
	//check friendship between invitedBy and invited user
	//insert into group_invites with status = "approved" if admin is invite
	//insert into group_invites with status = "pending" if not admin is invite
	//if admin is invite -> Add to conversation_members.

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(ctx, conversationID, userID)

	if err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	for _, iuser := range dto.InvitedUserId {
		if !s.fRepo.AlreadyFriends(ctx, userID, iuser) {
			continue
		}

		status := models.GroupPending
		if isGroupAdmin {
			status = models.GroupApproved
		}

		groupInvite := &models.GroupInvite{
			ConversationID: conversationID,
			InvitedBy:      userID,
			InvitedUserId:  iuser,
			Status:         status,
		}

		if err := s.giRepo.CreateGroupInvite(groupInvite); err != nil {
			return &e.NotFoundError{Message: err.Error()}
		}

		if isGroupAdmin {
			conversationMember := &models.ConversationMember{
				ConversationID: conversationID,
				UserID:         iuser,
			}
			if err := s.cmRepo.CreateConversationMember(conversationMember); err != nil {
				return &e.InternalServerError{Message: err.Error()}
			}
		}

	}

	return nil
}

func (s *conService) ModerateGroupInvite(
	ctx context.Context, conversationID string, inviteID string, userID string, dto dto.ModerateGroupInviteDto,
) error {

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(ctx, conversationID, userID)

	if err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if !isGroupAdmin {
		return &e.ForbiddenError{Message: "You are not an admin of this group"}
	}

	mgi := &models.GroupInvite{
		InvitedUserId:  inviteID,
		ConversationID: conversationID,
		Status:         dto.Status,
	}

	if err := s.giRepo.ModerateGroupInvite(mgi); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	conversationMember := &models.ConversationMember{
		ConversationID: conversationID,
		UserID:         inviteID,
	}

	if err := s.cmRepo.CreateConversationMember(conversationMember); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, please try again later"}
	}

	return nil
}

func (s *conService) AssignAdmin(ctx context.Context, conversationID string, userID string, dto dto.AssignAdminDto) error {

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(ctx, conversationID, userID)

	if err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if !isGroupAdmin {
		return &e.ForbiddenError{Message: "You are not an admin of this group"}
	}

	for _, iuser := range dto.AssignUserID {
		groupAdmin := &models.GroupAdmin{
			ConversationID: conversationID,
			UserID:         iuser,
		}

		if err := s.gaRepo.CreateGroupAdmin(groupAdmin); err != nil {
			return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
		}
	}

	return nil
}

func (s *conService) GetGroupMembers(ctx context.Context, conversationId string) ([]models.User, error) {
	users, err := s.cRepo.GetGroupMembers(ctx, conversationId)
	if err != nil {
		return nil, &e.InternalServerError{Message: err.Error()}
	}
	return users, nil
}

func (s *conService) GetGroupsById(ctx context.Context, userID string) ([]models.Conversation, error) {
	conversations, err := s.cRepo.GetGroupsById(ctx, userID)

	if err != nil {
		return nil, &e.InternalServerError{Message: err.Error()}
	}

	return conversations, nil
}
