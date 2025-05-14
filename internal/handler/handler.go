package handler

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/services"
	"github.com/aungsannphyo/ywartalk/internal/store"
)

type HandlerSet struct {
	UserHandler          *UserHandler
	FriendRequestHandler *FriendRequestHandler
	FriendHandler        *FriendHandler
	MessageHandler       *MessageHandler
	ConversationsHandler *ConversationsHandler
}

func InitHandler(db *sql.DB) *HandlerSet {
	//Repositories
	repoFactory := store.NewRepositoryFactory(db)
	serviceFactory := services.NewServiceFactory(repoFactory)
	//Services

	return &HandlerSet{
		UserHandler:          NewUserHandler(serviceFactory.UserService()),
		FriendRequestHandler: NewFriendRequestHandler(serviceFactory.FriendRequestService()),
		FriendHandler:        NewFriendHandler(serviceFactory.FriendService()),
		MessageHandler:       NewMessageHandler(serviceFactory.MessageService()),
		ConversationsHandler: NewConversationHandler(serviceFactory.ConversationService()),
	}
}
