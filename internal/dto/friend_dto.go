package dto

import (
	"strings"

	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type CreateFriendDto struct {
	UserID   string `json:"userId"`
	FriendID string `json:"friendId"`
}

type UnFriendDto struct {
	UserID   string `json:"userId"`
	FriendID string `json:"friendId"`
}

func ValidateUnFriendRequest(uf UnFriendDto) error {
	var errs v.ValidationErrors

	if strings.TrimSpace(uf.UserID) == "" {
		errs = append(errs, v.ValidationError{Field: "User Id", Message: "User Id is required"})
	}

	if strings.TrimSpace(uf.FriendID) == "" {
		errs = append(errs, v.ValidationError{Field: "Friend Request Id", Message: "Friend Request Id is required"})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
