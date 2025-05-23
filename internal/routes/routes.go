package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.HandlerSet) {
	api := r.Group("/api")
	api.Static("/static/profiles", "./uploads/profiles")
	AuthRoutes(api.Group("/"), h)
	RegisterUserRoutes(api.Group("/user"), h)
	RegisterFriendRequestRoutes(api.Group("/friend"), h)
	RegisterFriendRoutes(api.Group("/friend"), h)
	RegisterConversationRoutes(api.Group("/group"), h)
	RegisterWebSocketRoute(api.Group("/"), h)
	RegisterMessageRoute(api.Group("/message"), h)
}
