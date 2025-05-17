package dto

import (
	"log"
	"strings"

	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type SendPrivateMessageDto struct {
	ReceiverId string `json:"receiver_id"`
	Content    string `json:"content"`
}

type SendGroupMessageDto struct {
	Content string `json:"content"`
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
