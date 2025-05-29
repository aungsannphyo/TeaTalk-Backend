package models

import (
	"time"
)

type User struct {
	ID               string    `json:"id"`
	UserIdentity     string    `json:"user_identity"`
	Email            string    `json:"email"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	CreatedAt        time.Time `json:"created_at"`
	Salt             []byte    `json:"salt"`
	EncryptedUserKey []byte    `json:"encrypted_user_key"`
	UserKeyNonce     []byte    `json:"user_key_nonce"`
}
