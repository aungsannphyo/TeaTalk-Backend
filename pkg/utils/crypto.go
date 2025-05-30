package utils

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

const (
	KeyLength  = 32
	PBKDF2Iter = 100_000
)

func DeriveKey(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, PBKDF2Iter, KeyLength, sha256.New)
}
