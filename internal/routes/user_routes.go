package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.GET("/:id", h.UserHandler.GetUser)
}
