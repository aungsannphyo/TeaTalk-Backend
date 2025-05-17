package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterFriendRequestRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.POST("/request", h.FriendRequestHandler.SendFriendRequestHandler)
	rg.POST("/decide-request", h.FriendRequestHandler.DecideFriendRequestHandler)
}
