package handler

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/services"
	store "github.com/aungsannphyo/ywartalk/internal/store/repo"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/internal/websocket/group"
	"github.com/aungsannphyo/ywartalk/internal/websocket/private"
)

type HandlerSet struct {
	UserHandler          *UserHandler
	FriendRequestHandler *FriendRequestHandler
	FriendHandler        *FriendHandler
	ConversationsHandler *ConversationsHandler
	PrivateHubHandler    *private.WebSocketPrivateHandler
	GroupHubHandler      *group.WebSocketGroupHandler
}

func InitHandler(db *sql.DB) *HandlerSet {

	//sql loader
	sqlLoader := sqlloader.SQLLoaderFactory(sqlloader.EmbedLoader)

	//Repositories
	repoFactory := store.NewRepositoryFactory(db, sqlLoader)
	//Services
	serviceFactory := services.NewServiceFactory(repoFactory)

	//WebSocket
	privateHub := private.NewPrivateHub()
	groupHub := group.NewGroupHub(serviceFactory.ConversationService())

	privateHubHandler := private.NewWebSocketPrivateHandler(
		privateHub,
		serviceFactory.MessageService(),
		serviceFactory.UserService(),
	)

	groupHubHandler := group.NewWebSocketGroupHandler(
		groupHub,
		serviceFactory.MessageService(),
	)

	go privateHub.RunPrivateWebSocket()
	go groupHub.RunGroupWebSocket()

	return &HandlerSet{
		PrivateHubHandler:    privateHubHandler,
		GroupHubHandler:      groupHubHandler,
		UserHandler:          NewUserHandler(serviceFactory.UserService()),
		FriendRequestHandler: NewFriendRequestHandler(serviceFactory.FriendRequestService()),
		FriendHandler:        NewFriendHandler(serviceFactory.FriendService()),
		ConversationsHandler: NewConversationHandler(serviceFactory.ConversationService()),
	}
}
