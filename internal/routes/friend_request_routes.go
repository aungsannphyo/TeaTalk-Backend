package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterFriendRequestRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.POST("/requests", h.FriendRequestHandler.SendFriendRequestHandler)
	rg.POST("/decide-request", h.FriendRequestHandler.DecideFriendRequestHandler)
	rg.GET("/requests", h.FriendRequestHandler.GetAllFriendRequestHandler)
}
