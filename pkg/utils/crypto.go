package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

const (
	KeyLength   = 32
	PBKDF2Iter  = 100_000
	SaltLength  = 16
	NonceLength = 12
)

// Generate random bytes
func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	return bytes, err
}

// Derive key using PBKDF2
func DeriveKey(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, PBKDF2Iter, KeyLength, sha256.New)
}

// Encrypt with AES-GCM
func Encrypt(plainText, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	nonce, _ := GenerateRandomBytes(NonceLength)
	aesgcm, _ := cipher.NewGCM(block)
	cipherText := aesgcm.Seal(nil, nonce, plainText, nil)
	return cipherText, nonce, nil
}

// Decrypt with AES-GCM
func Decrypt(cipherText, nonce, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, _ := cipher.NewGCM(block)
	plain, err := aesgcm.Open(nil, nonce, cipherText, nil)
	return plain, err
}
