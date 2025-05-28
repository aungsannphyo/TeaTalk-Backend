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

			// Send the current list of online users only to the new client
			h.sendOnlineUsersToClient(c)
			go c.onlineManager.SetUserOnline(c.userID)
			// Broadcast this user's status online to all clients (including the new client)
			h.broadcastStatusToAll(c.userID, 1) // 1 for online status

			log.Printf("User %s connected", c.userID)

		case c := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[c.userID]; ok {
				delete(h.clients, c.userID)
				close(c.send)
			}
			h.mu.Unlock()

			h.broadcastStatusToAll(c.userID, 0)
			go c.onlineManager.SetUserOffline(c.userID)

			// 0 for offline status

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

func (h *PrivateHub) sendOnlineUsersToClient(c *PrivateClient) {
	onlineUsers := c.onlineManager.GetOnlineUsers()
	statuses := make([]StatusUpdate, 0)
	for _, uid := range onlineUsers {
		if uid == c.userID {
			continue
		}
		statuses = append(statuses, StatusUpdate{
			UserID: uid,
			Status: 1,
		})
	}

	if len(statuses) > 0 {
		if data, err := json.Marshal(statuses); err == nil {
			c.send <- data
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
