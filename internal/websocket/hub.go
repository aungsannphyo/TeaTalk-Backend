package websocket

import (
	"log"
	"sync"
)

type Hub struct {
	clients        map[string]*Client
	groups         map[string]map[string]*Client
	register       chan *Client
	unregister     chan *Client
	privateMessage chan PrivateMessage
	groupMessage   chan GroupMessage
	mu             sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:        make(map[string]*Client),
		groups:         make(map[string]map[string]*Client),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		privateMessage: make(chan PrivateMessage),
		groupMessage:   make(chan GroupMessage),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.mu.Lock()
			h.clients[c.UserID] = c
			h.mu.Unlock()
			log.Printf("User %s connected", c.UserID)

		case c := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[c.UserID]; ok {
				delete(h.clients, c.UserID)
				close(c.Send)
				for groupID := range h.groups {
					h.RemoveUserFromGroup(groupID, c.UserID)
				}
				log.Printf("User %s disconnected", c.UserID)
			}
			h.mu.Unlock()

		case pm := <-h.privateMessage:
			h.mu.RLock()
			recipient, ok := h.clients[pm.ToUserID]
			h.mu.RUnlock()
			if ok {
				recipient.Send <- pm.Content
			}

		case gm := <-h.groupMessage:
			h.mu.RLock()
			members, ok := h.groups[gm.GroupID]
			h.mu.RUnlock()
			if ok {
				for uid, member := range members {
					if uid != gm.FromUserID {
						member.Send <- gm.Content
					}
				}
			}
		}
	}
}

func (h *Hub) AddUserToGroup(groupID, userID string, c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.groups[groupID] == nil {
		h.groups[groupID] = make(map[string]*Client)
	}
	h.groups[groupID][userID] = c
}

func (h *Hub) RemoveUserFromGroup(groupID, userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if members, ok := h.groups[groupID]; ok {
		delete(members, userID)
		if len(members) == 0 {
			delete(h.groups, groupID)
		}
	}
}
