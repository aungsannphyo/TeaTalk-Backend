package dto

import (
	"strings"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type SendFriendRequestDto struct {
	SenderId   string `json:"senderId"`
	ReceiverId string `json:"receiverId"`
}

type DecideFriendRequestDto struct {
	UserId          string                     `json:"userId"`
	Status          models.FriendRequestStatus `json:"status"`
	FriendRequestId string                     `json:"friendRequestId"`
}

func ValidateSendFriendRequest(sfrDto SendFriendRequestDto) error {
	var errs v.ValidationErrors
	if strings.TrimSpace(sfrDto.SenderId) == "" {
		errs = append(errs, v.ValidationError{Field: "Sender Id", Message: "Sender Id is required"})
	}
	if strings.TrimSpace(sfrDto.ReceiverId) == "" {
		errs = append(errs, v.ValidationError{Field: "Receiver Id", Message: "Receiver Id is required"})
	}

	return nil
}

func ValidateDecideFriendRequest(dfr DecideFriendRequestDto) error {
	var errs v.ValidationErrors

	if strings.TrimSpace(dfr.UserId) == "" {
		errs = append(errs, v.ValidationError{Field: "User Id", Message: "User Id is required"})
	}

	if strings.TrimSpace(dfr.FriendRequestId) == "" {
		errs = append(errs, v.ValidationError{Field: "Friend Request Id", Message: "Friend Request Id is required"})
	}

	if dfr.Status != models.StatusAccepted && dfr.Status != models.StatusRejected {
		errs = append(errs, v.ValidationError{
			Field:   "Status",
			Message: "Status should be ACCEPTED or REJECTED",
		})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
