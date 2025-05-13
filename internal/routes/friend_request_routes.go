package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterFriendRequestRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.POST("/requests", h.FriendRequestHandler.SendFriendRequest)
	rg.PATCH("/decide-request", h.FriendRequestHandler.DecideFriendRequest)
}
