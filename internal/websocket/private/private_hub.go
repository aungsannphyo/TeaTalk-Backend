package private

import (
	"log"
	"sync"
)

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

			log.Printf("User %s connected", c.userID)

		case c := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[c.userID]; ok {
				delete(h.clients, c.userID)
				close(c.send)
			}
			h.mu.Unlock()

			go c.onlineManager.SetUserOffline(c.userID)

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
