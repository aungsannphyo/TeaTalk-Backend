package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMessageRoute(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.GET("/", h.MessageHandler.GetMessages)
}
