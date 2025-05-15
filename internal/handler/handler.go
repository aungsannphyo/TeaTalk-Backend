package handler

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/services"
	"github.com/aungsannphyo/ywartalk/internal/store"
	"github.com/aungsannphyo/ywartalk/internal/websocket"
	"github.com/gin-gonic/gin"
)

type HandlerSet struct {
	UserHandler          *UserHandler
	FriendRequestHandler *FriendRequestHandler
	FriendHandler        *FriendHandler
	ConversationsHandler *ConversationsHandler
	Hub                  *websocket.Hub
}

func InitHandler(db *sql.DB) *HandlerSet {
	//WebSocket
	hub := websocket.NewHub()
	go hub.Run()

	//Repositories
	repoFactory := store.NewRepositoryFactory(db)
	//Services
	serviceFactory := services.NewServiceFactory(repoFactory)

	return &HandlerSet{
		Hub:                  hub,
		UserHandler:          NewUserHandler(serviceFactory.UserService()),
		FriendRequestHandler: NewFriendRequestHandler(serviceFactory.FriendRequestService()),
		FriendHandler:        NewFriendHandler(serviceFactory.FriendService()),
		ConversationsHandler: NewConversationHandler(serviceFactory.ConversationService()),
	}
}

func (h *HandlerSet) WebSocketHandler(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// Upgrade HTTP to WS connection
	websocket.HandleWebSocket(c.Writer, c.Request, userID.(string), h.Hub)
}
