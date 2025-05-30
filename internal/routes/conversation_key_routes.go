package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterConversationKeyRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.GET("/:conversationId/user/:userId", h.ConversationKeyHandler.GetConversationKey)
	rg.POST("/create-conversation-key", h.ConversationKeyHandler.CreateConversationKey)

}
