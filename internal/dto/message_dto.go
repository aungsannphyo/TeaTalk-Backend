package dto

import (
	"log"
	"strings"
	"time"

	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type SendPrivateMessageDto struct {
	ReceiverId string `json:"receiverId"`
	Content    string `json:"content"`
}

type SendGroupMessageDto struct {
	Content string `json:"content"`
}

type MessagesDto struct {
	MessageID        string    `json:"messageId"`
	SenderID         string    `json:"senderId"`
	MemberID         string    `json:"memberId"`
	Content          string    `json:"content"`
	IsRead           bool      `json:"isRead"`
	SeenByName       *string   `json:"seenByName,omitempty"`
	MessageCreatedAt time.Time `json:"messageCreatedAt"`
}

func ValidateSendMessageRequest(smDto SendPrivateMessageDto) error {
	var errs v.ValidationErrors
	if strings.TrimSpace(smDto.ReceiverId) == "" {
		log.Print("Check Reciver")
		errs = append(errs, v.ValidationError{Field: "Receiver Id", Message: "Receiver Id is required"})
	}

	if strings.TrimSpace(smDto.Content) == "" {
		log.Print("Check Content")
		errs = append(errs, v.ValidationError{Field: "Content", Message: "Content is required"})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func ValidateSendGroupMessage(sgm SendGroupMessageDto) error {
	var errs v.ValidationErrors

	if strings.TrimSpace(sgm.Content) == "" {
		log.Print("Check Content")
		errs = append(errs, v.ValidationError{Field: "Content", Message: "Content is required"})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
