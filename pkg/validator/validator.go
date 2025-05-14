package validator

import (
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	msg := "Validation failed:\n"

	for _, err := range v {
		msg += fmt.Sprintf("  - %s: %s\n", err.Field, err.Message)
	}
	return msg
}

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
