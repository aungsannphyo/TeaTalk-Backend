package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.POST("/register", h.UserHandler.RegisterHandler)
	rg.POST("/login", h.UserHandler.LoginHandler)
}
