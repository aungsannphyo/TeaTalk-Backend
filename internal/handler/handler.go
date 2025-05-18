package handler

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/services"
	store "github.com/aungsannphyo/ywartalk/internal/store/repo"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/internal/websocket"
	ws "github.com/aungsannphyo/ywartalk/internal/websocket"
)

type HandlerSet struct {
	UserHandler          *UserHandler
	FriendRequestHandler *FriendRequestHandler
	FriendHandler        *FriendHandler
	ConversationsHandler *ConversationsHandler
	HubHandler           *ws.WebSocketHandler
}

func InitHandler(db *sql.DB) *HandlerSet {

	//sql loader
	sqlLoader := sqlloader.SQLLoaderFactory(sqlloader.EmbedLoader)

	//Repositories
	repoFactory := store.NewRepositoryFactory(db, sqlLoader)
	//Services
	serviceFactory := services.NewServiceFactory(repoFactory)

	//WebSocket
	hub := websocket.NewHub(serviceFactory.ConversationService())
	go hub.Run()

	websocketHandler := ws.NewWebSocketHandler(
		hub,
		serviceFactory.MessageService(),
		serviceFactory.UserService())

	return &HandlerSet{
		HubHandler:           websocketHandler,
		UserHandler:          NewUserHandler(serviceFactory.UserService()),
		FriendRequestHandler: NewFriendRequestHandler(serviceFactory.FriendRequestService()),
		FriendHandler:        NewFriendHandler(serviceFactory.FriendService()),
		ConversationsHandler: NewConversationHandler(serviceFactory.ConversationService()),
	}
}
