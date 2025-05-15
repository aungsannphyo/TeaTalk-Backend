package handler

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/services"
	"github.com/aungsannphyo/ywartalk/internal/store"
	"github.com/aungsannphyo/ywartalk/internal/websocket"
)

type HandlerSet struct {
	UserHandler          *UserHandler
	FriendRequestHandler *FriendRequestHandler
	FriendHandler        *FriendHandler
	ConversationsHandler *ConversationsHandler
	HubHandler           *WebSocketHandler
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
		HubHandler:           NewWebSocketHandler(hub),
		UserHandler:          NewUserHandler(serviceFactory.UserService()),
		FriendRequestHandler: NewFriendRequestHandler(serviceFactory.FriendRequestService()),
		FriendHandler:        NewFriendHandler(serviceFactory.FriendService()),
		ConversationsHandler: NewConversationHandler(serviceFactory.ConversationService()),
	}
}
