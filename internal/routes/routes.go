package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.HandlerSet) {
	api := r.Group("/api")
	api.Static("/static/profiles", "./uploads/profiles")
	AuthRoutes(api.Group("/"), h)
	RegisterUserRoutes(api.Group("/users"), h)
	RegisterFriendRequestRoutes(api.Group("/friends"), h)
	RegisterFriendRoutes(api.Group("/friends"), h)
	RegisterConversationRoutes(api.Group("/conversations"), h)
	RegisterWebSocketRoute(api.Group("/"), h)
	RegisterMessageRoute(api.Group("/messages"), h)
}
