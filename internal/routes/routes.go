package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.HandlerSet) {
	api := r.Group("/api")
	AuthRoutes(api.Group("/"), h)
	RegisterUserRoutes(api.Group("/user"), h)
	RegisterFriendRequestRoutes(api.Group("/friend"), h)
	RegisterFriendRoutes(api.Group("/friend"), h)
}
