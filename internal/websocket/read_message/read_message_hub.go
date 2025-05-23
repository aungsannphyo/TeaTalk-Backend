package readmessage

import (
	"log"
	"sync"
)

type ReadMessageHub struct {
	register   chan *ReadMessageClient
	unregister chan *ReadMessageClient
	clients    map[string]*ReadMessageClient
	broadcast  chan WSReadMessage
	mu         sync.RWMutex
}

func NewReadMessageHub() *ReadMessageHub {
	return &ReadMessageHub{
		clients:    make(map[string]*ReadMessageClient),
		register:   make(chan *ReadMessageClient),
		unregister: make(chan *ReadMessageClient),
		broadcast:  make(chan WSReadMessage),
	}
}

func (h *ReadMessageHub) RunReadMessageHub() {
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

			log.Printf("User %s connected", c.userID)

		case c := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[c.userID]; ok {
				delete(h.clients, c.userID)
				close(c.send)
			}
			h.mu.Unlock()

		case pm := <-h.broadcast:
			h.mu.RLock()
			recipient, ok := h.clients[pm.ReaderId]
			h.mu.RUnlock()
			if ok {
				recipient.send <- pm.MessageID
			}
		}
	}
}
