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

func (s *conService) CreateGroup(userID string, dto dto.CreateGroupDto) error {
	cID := uuid.NewString()

	conversation := &models.Conversation{
		ID:        cID,
		IsGroup:   true,
		Name:      &dto.Name,
		CreatedBy: &userID,
	}

	conversationMember := &models.ConversationMember{
		ConversationID: cID,
		UserID:         userID,
	}

	groupAdmin := &models.GroupAdmin{
		ConversationID: cID,
		UserID:         userID,
	}

	if err := s.cRepo.CreateConversation(conversation); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if err := s.cmRepo.CreateConversationMember(conversationMember); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if err := s.gaRepo.CreateGroupAdmin(groupAdmin); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}
	return nil
}

func (s *conService) UpdateGroupName(groupID string, dto dto.UpdateGroupNameDto) error {

	uc := &models.Conversation{
		ID:   groupID,
		Name: &dto.Name,
	}

	if err := s.cRepo.UpdateGroupName(uc); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}
	return nil
}

func (s *conService) InviteGroup(ctx context.Context, groupID string, userID string, dto dto.InviteGroupDto) error {
	//need to check current user is group admin or not
	//check friendship between invitedBy and invited user
	//insert into group_invites with status = "approved" if admin is invite
	//insert into group_invites with status = "pending" if not admin is invite
	//if admin is invite -> Add to conversation_members.

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(ctx, groupID, userID)

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
			ConversationID: groupID,
			InvitedBy:      userID,
			InvitedUserId:  iuser,
			Status:         status,
		}

		if err := s.giRepo.CreateGroupInvite(groupInvite); err != nil {
			return &e.NotFoundError{Message: err.Error()}
		}

		if isGroupAdmin {
			conversationMember := &models.ConversationMember{
				ConversationID: groupID,
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
	ctx context.Context, groupID string, inviteID string, userID string, dto dto.ModerateGroupInviteDto,
) error {

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(ctx, groupID, userID)

	if err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if !isGroupAdmin {
		return &e.ForbiddenError{Message: "You are not an admin of this group"}
	}

	mgi := &models.GroupInvite{
		InvitedUserId:  inviteID,
		ConversationID: groupID,
		Status:         dto.Status,
	}

	if err := s.giRepo.ModerateGroupInvite(mgi); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	conversationMember := &models.ConversationMember{
		ConversationID: groupID,
		UserID:         inviteID,
	}

	if err := s.cmRepo.CreateConversationMember(conversationMember); err != nil {
		return &e.InternalServerError{Message: "Something went wrong, please try again later"}
	}

	return nil
}

func (s *conService) AssignAdmin(ctx context.Context, groupID string, userID string, dto dto.AssignAdminDto) error {

	isGroupAdmin, err := s.gaRepo.IsGroupAdmin(ctx, groupID, userID)

	if err != nil {
		return &e.InternalServerError{Message: "Something went wrong, Please try again later"}
	}

	if !isGroupAdmin {
		return &e.ForbiddenError{Message: "You are not an admin of this group"}
	}

	for _, iuser := range dto.AssignUserID {
		groupAdmin := &models.GroupAdmin{
			ConversationID: groupID,
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
