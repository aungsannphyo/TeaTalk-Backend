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

// EncryptConversationKey encrypts the conversation key with a user's key
func EncryptConversationKey(convKey, userKey []byte) (encryptedKey []byte, nonce []byte, err error) {
	return Encrypt(convKey, userKey)
}

// DecryptConversationKey decrypts the conversation key with a user's key
func DecryptConversationKey(encryptedKey, nonce, userKey []byte) ([]byte, error) {
	return Decrypt(encryptedKey, nonce, userKey)
}

// EncryptMessageContent encrypts a message content with the conversation key
func EncryptMessageContent(plainMessage, convKey []byte) (encryptedMessage []byte, nonce []byte, err error) {
	return Encrypt(plainMessage, convKey)
}

// DecryptMessageContent decrypts a message content with the conversation key
func DecryptMessageContent(encryptedMessage, nonce, convKey []byte) ([]byte, error) {
	return Decrypt(encryptedMessage, nonce, convKey)
}

// EncryptUserKey encrypts the conversation key with a user's key
func EncryptUserKey(convKey, userKey []byte) (encryptedKey []byte, nonce []byte, err error) {
	return Encrypt(convKey, userKey)
}

// DecryptUserKey decrypts the conversation key with a user's key
func DecryptUserKey(encryptedKey, nonce, userKey []byte) ([]byte, error) {
	return Decrypt(encryptedKey, nonce, userKey)
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
