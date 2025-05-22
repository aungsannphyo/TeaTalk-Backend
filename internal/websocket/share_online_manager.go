package websocket

import (
	"log"
	"sync"

	"github.com/aungsannphyo/ywartalk/internal/domain/service"
)

type SharedOnlineManager struct {
	mu          sync.Mutex
	onlineUsers map[string]int // userID -> connection count
	userService service.UserService
}

func NewSharedOnlineManager(userService service.UserService) *SharedOnlineManager {
	return &SharedOnlineManager{
		onlineUsers: make(map[string]int),
		userService: userService,
	}
}

func (m *SharedOnlineManager) SetUserOnline(userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	log.Println("USERID", userID)

	m.onlineUsers[userID]++
	if m.onlineUsers[userID] == 1 {
		if err := m.userService.SetUserOnline(userID); err != nil {
			log.Printf("Failed to set user %s online: ", userID)
		}
	}
}

func (m *SharedOnlineManager) SetUserOffline(userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	log.Println("USERID", userID)

	if m.onlineUsers[userID] > 0 {
		m.onlineUsers[userID]--
		if m.onlineUsers[userID] == 0 {
			// Last connection gone
			delete(m.onlineUsers, userID)
			if err := m.userService.SetUserOffline(userID); err != nil {
				log.Printf("Failed to set user %s offline: ", userID)
			}
		}
	}
}
