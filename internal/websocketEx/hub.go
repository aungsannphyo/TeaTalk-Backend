package websocketEx

import (
	"log"
	"sync"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type Hub struct {
	sync.RWMutex

	clients map[*Client]bool

	broadcast  chan *models.Message
	register   chan *Client
	unregister chan *Client

	messages []*models.Message
}

func NewHub() *Hub {
	return &Hub{
		clients:    map[*Client]bool{},
		broadcast:  make(chan *models.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()

			log.Printf("client registered %s", client.id)

		case client := <-h.unregister:
			h.Lock()
			if _, ok := h.clients[client]; ok {
				close(client.send)
				log.Printf("client unregistered %s", client.id)
				delete(h.clients, client)
			}
			h.Unlock()
		case msg := <-h.broadcast:
			h.RLock()
			h.messages = append(h.messages, msg)
			h.RUnlock()
		}
	}
}
