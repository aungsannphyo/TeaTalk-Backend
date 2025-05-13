package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterFriendRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.DELETE("/unfriend", h.FriendHandler.MakeUnFriend)
}
