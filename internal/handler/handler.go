package handler

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/service"
	"github.com/aungsannphyo/ywartalk/internal/store"
)

type HandlerSet struct {
	UserHandler          *UserHandler
	FriendRequestHandler *FriendRequestHandler
	FriendHandler        *FriendHandler
}

func InitHandler(db *sql.DB) *HandlerSet {
	//Repositories
	userRepo := store.NewUserRepo(db)
	friendRequestRepo := store.NewFriendRequestRepo(db)
	friendRepo := store.NewFriendRepo(db)
	friendRequestLogRepo := store.NewFriendRequestLogRepo(db)

	//Services
	userService := service.NewUserService(userRepo)
	friendRequestService := service.NewFriendRequestService(
		friendRequestRepo,
		friendRepo,
		friendRequestLogRepo)
	friendService := service.NewFriendService(
		friendRepo,
		friendRequestRepo,
		friendRequestLogRepo,
	)

	return &HandlerSet{
		UserHandler:          NewUserHandler(userService),
		FriendRequestHandler: NewFriendRequestHandler(friendRequestService),
		FriendHandler:        NewFriendHandler(friendService),
	}
}
