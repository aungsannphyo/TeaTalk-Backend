package dto

import (
	"strings"

	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type CreateGroupDto struct {
	IsGroup bool   `json:"isGroup"`
	Name    string `json:"name"`
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
