package dto

import (
	"strings"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type CreateGroupDto struct {
	IsGroup bool   `json:"isGroup"`
	Name    string `json:"name"`
}

type UpdateGroupNameDto struct {
	Name string `json:"name"`
}

type InviteGroupDto struct {
	InvitedUserId []string `json:"invitedUserId"`
}

type AssignAdminDto struct {
	InvitedUserId []string `json:"invitedUserId"`
}

type ModerateGroupInviteDto struct {
	Status models.GroupInviteStatus `json:"status"`
}

func ValidateModerateGroupInvite(mgi ModerateGroupInviteDto) error {
	var errs v.ValidationErrors

	if mgi.Status != models.GroupApproved && mgi.Status != models.GroupRejected {
		errs = append(errs, v.ValidationError{
			Field:   "Status",
			Message: "Status should be APPROVED or REJECTED"})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func ValidateCreateGroup(g CreateGroupDto) error {
	var errs v.ValidationErrors

	if strings.TrimSpace(g.Name) == "" {
		errs = append(errs, v.ValidationError{Field: "Name", Message: "Friend Request Id is required"})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func ValidateInviteGroup(ig InviteGroupDto) error {
	var errs v.ValidationErrors

	if len(ig.InvitedUserId) == 0 {
		errs = append(errs, v.ValidationError{Field: "Invite Users Id", Message: "Invite users list is empty! "})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func ValidateAssignAdmin(aa AssignAdminDto) error {
	var errs v.ValidationErrors

	if len(aa.InvitedUserId) == 0 {
		errs = append(errs, v.ValidationError{Field: "Invite Users Id", Message: "Invite users list is empty! "})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
