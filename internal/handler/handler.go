package handler

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/services"
	store "github.com/aungsannphyo/ywartalk/internal/store/repo"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	ws "github.com/aungsannphyo/ywartalk/internal/websocket"
	"github.com/aungsannphyo/ywartalk/internal/websocket/group"
	"github.com/aungsannphyo/ywartalk/internal/websocket/private"
	readmessage "github.com/aungsannphyo/ywartalk/internal/websocket/read_message"
)

type HandlerSet struct {
	UserHandler          *UserHandler
	FriendRequestHandler *FriendRequestHandler
	FriendHandler        *FriendHandler
	ConversationsHandler *ConversationsHandler
	PrivateHubHandler    *private.WebSocketPrivateHandler
	GroupHubHandler      *group.WebSocketGroupHandler
	MessageHandler       *MessageHandler
	ReadMessageHandler   *readmessage.WebSocketReadMessageHandler
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
	readMessageHub := readmessage.NewReadMessageHub()

	onlineManager := ws.NewSharedOnlineManager(serviceFactory.UserService())

	privateHubHandler := private.NewWebSocketPrivateHandler(
		privateHub,
		serviceFactory.MessageService(),
		onlineManager,
	)

	groupHubHandler := group.NewWebSocketGroupHandler(
		groupHub,
		serviceFactory.MessageService(),
		onlineManager,
	)

	msgReadHubHandler := readmessage.NewWebSocketReadMessageHandler(
		readMessageHub,
		serviceFactory.MessageReadService(),
	)

	go privateHub.RunPrivateWebSocket()
	go groupHub.RunGroupWebSocket()
	go readMessageHub.RunReadMessageHub()

	return &HandlerSet{
		PrivateHubHandler:    privateHubHandler,
		GroupHubHandler:      groupHubHandler,
		ReadMessageHandler:   msgReadHubHandler,
		UserHandler:          NewUserHandler(serviceFactory.UserService()),
		FriendRequestHandler: NewFriendRequestHandler(serviceFactory.FriendRequestService()),
		FriendHandler:        NewFriendHandler(serviceFactory.FriendService()),
		ConversationsHandler: NewConversationHandler(serviceFactory.ConversationService()),
		MessageHandler:       NewMessageHandler(serviceFactory.MessageService()),
	}
}
