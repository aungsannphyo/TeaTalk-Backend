package group

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/aungsannphyo/ywartalk/internal/domain/service"
)

type StatusUpdate struct {
	UserID string `json:"userId"`
	Status int    `json:"status"`
}

type GroupHub struct {
	clients     map[string]*GroupClient
	groups      map[string]map[string]*GroupClient
	register    chan *GroupClient
	unregister  chan *GroupClient
	broadcast   chan GroupMessage
	mu          sync.RWMutex
	convService service.ConversationService
}

func NewGroupHub(convService service.ConversationService) *GroupHub {
	return &GroupHub{
		clients:     make(map[string]*GroupClient),
		groups:      make(map[string]map[string]*GroupClient),
		register:    make(chan *GroupClient),
		unregister:  make(chan *GroupClient),
		broadcast:   make(chan GroupMessage),
		convService: convService,
	}
}

func (h *GroupHub) RunGroupWebSocket() {
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

			go c.onlineManager.SetUserOnline(c.userID)

			ctx := context.Background()
			groups, err := h.convService.GetGroupsById(ctx, c.userID)

			if err != nil {
				log.Println("fetch error")
			}

			for _, group := range groups {
				h.AddUserToGroup(group.ID, c.userID, c)
			}

			h.broadcastStatusToGroups(c.userID, 1) // 1 for online status

		case c := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[c.userID]; ok {
				delete(h.clients, c.userID)
				close(c.send)
				for groupID := range h.groups {
					h.RemoveUserFromGroup(groupID, c.userID)
				}
				log.Printf("User %s disconnected", c.userID)
			}
			h.mu.Unlock()

			go c.onlineManager.SetUserOffline(c.userID)

			h.broadcastStatusToGroups(c.userID, 0) // 0 for offline status

		case gm := <-h.broadcast:
			h.mu.RLock()
			members, ok := h.groups[gm.GroupID]
			h.mu.RUnlock()
			if ok {
				for uid, member := range members {
					if uid != gm.SenderID {
						member.send <- gm.Content
					}
				}
			}
		}
	}
}

func (h *GroupHub) AddUserToGroup(groupID, userID string, c *GroupClient) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.groups[groupID] == nil {
		h.groups[groupID] = make(map[string]*GroupClient)
	}
	h.groups[groupID][userID] = c

}

func (h *GroupHub) RemoveUserFromGroup(groupID, userID string) {
	if members, ok := h.groups[groupID]; ok {
		delete(members, userID)
		if len(members) == 0 {
			delete(h.groups, groupID)
		}
	}
}

func (h *GroupHub) broadcastStatusToGroups(userID string, status int) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	update, _ := json.Marshal(StatusUpdate{UserID: userID, Status: status})
	for _, members := range h.groups {
		if _, ok := members[userID]; ok {
			for uid, member := range members {
				if uid != userID {
					select {
					case member.send <- update:
					default:
					}
				}
			}
		}
	}
}
