package private

import (
	"encoding/json"
	"log"
	"sync"
)

type StatusUpdate struct {
	UserID string `json:"userId"`
	Status int    `json:"status"`
}

type PrivateHub struct {
	register   chan *PrivateClient
	unregister chan *PrivateClient
	clients    map[string]*PrivateClient
	broadcast  chan PrivateMessage
	mu         sync.RWMutex
}

func NewPrivateHub() *PrivateHub {
	return &PrivateHub{
		clients:    make(map[string]*PrivateClient),
		register:   make(chan *PrivateClient),
		unregister: make(chan *PrivateClient),
		broadcast:  make(chan PrivateMessage),
	}
}

func (h *PrivateHub) RunPrivateWebSocket() {
	for {
		select {
		case c := <-h.register:
			h.mu.Lock()

			// Close old connection if exists
			if existing, ok := h.clients[c.userID]; ok {
				existing.conn.Close()
			}

			h.clients[c.userID] = c
			h.mu.Unlock()

			go c.onlineManager.SetUserOnline(c.userID)

			h.broadcastStatusToAll(c.userID, 1) // 1 for online status

			log.Printf("User %s connected", c.userID)

		case c := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[c.userID]; ok {
				delete(h.clients, c.userID)
				close(c.send)
			}
			h.mu.Unlock()

			go c.onlineManager.SetUserOffline(c.userID)

			h.broadcastStatusToAll(c.userID, 0) // 0 for offline status

		case pm := <-h.broadcast:
			h.mu.RLock()
			recipient, ok := h.clients[pm.ReceiverID]
			h.mu.RUnlock()
			if ok {
				recipient.send <- pm.Content
			}
		}
	}
}

func (h *PrivateHub) broadcastStatusToAll(userID string, status int) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	update, _ := json.Marshal(StatusUpdate{UserID: userID, Status: status})
	for _, client := range h.clients {
		select {
		case client.send <- update:
		default:
		}
	}
}
