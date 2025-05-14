package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMessageRoute(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.POST("/private", h.MessageHandler.SendPrivateMessage)
	rg.POST("/groups/:groupId", h.MessageHandler.SendGroupMessage)
}
