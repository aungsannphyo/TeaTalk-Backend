package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterWebSocketRoute(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.GET("/ws/private", h.PrivateHubHandler.WebSocketPrivateHandler)
	rg.GET("/ws/group", h.GroupHubHandler.NewWebSocketGroupHandler)
	rg.GET("/ws/read-message", h.ReadMessageHandler.WebSocketReadMessageHandler)
}
