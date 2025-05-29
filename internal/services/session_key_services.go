package services

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type SessionKeyCache struct {
	c *cache.Cache // userID -> []byte (actualUserKey)
}

func NewSessionKeyCache() *SessionKeyCache {
	// 15 min expiration, 30 min cleanup interval
	return &SessionKeyCache{
		c: cache.New(15*time.Minute, 30*time.Minute),
	}
}

func (s *SessionKeyCache) SetUserDecryptedKey(userID string, decryptedKey []byte) {
	s.c.Set(userID, decryptedKey, cache.DefaultExpiration)
}

func (s *SessionKeyCache) GetUserDecryptedKey(userID string) ([]byte, bool) {
	val, found := s.c.Get(userID)
	if !found {
		return nil, false
	}
	return val.([]byte), true
}

func (s *SessionKeyCache) DeleteUserKey(userID string) {
	s.c.Delete(userID)
}
