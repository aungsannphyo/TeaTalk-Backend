package dto

import (
	"strings"

	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type RegisterRequestDto struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Salt             string `json:"salt"`
	EncryptedUserKey string `json:"encryptedUserKey"`
	UserKeyNonce     string `json:"userKeyNonce"`
}

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateRegisterUser(user RegisterRequestDto) error {
	var errs v.ValidationErrors

	if strings.TrimSpace(user.Email) == "" {
		errs = append(errs, v.ValidationError{Field: "email", Message: "Email is required"})
	}
	if !v.IsValidEmail(user.Email) {
		errs = append(errs, v.ValidationError{Field: "email", Message: "Email format is invalid"})
	}
	if strings.TrimSpace(user.Username) == "" {
		errs = append(errs, v.ValidationError{Field: "username", Message: "Username is required"})
	}
	if len(user.Password) < 6 {
		errs = append(errs, v.ValidationError{Field: "password", Message: "Password must be at least 6 characters"})
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func ValidateLoginUser(user LoginRequestDto) error {
	var errs v.ValidationErrors

	if strings.TrimSpace(user.Email) == "" {
		errs = append(errs, v.ValidationError{Field: "email", Message: "Email is required"})
	}
	if !v.IsValidEmail(user.Email) {
		errs = append(errs, v.ValidationError{Field: "email", Message: "Email format is invalid"})
	}

	if strings.TrimSpace(user.Password) == "" {
		errs = append(errs, v.ValidationError{Field: "password", Message: "Password must be at least 6 characters"})
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
