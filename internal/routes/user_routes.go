package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.GET("/:userID", h.UserHandler.GetUserHandler)
	rg.GET("/:userID/chat-list", h.UserHandler.GetChatListByUserIdHandler)
	rg.POST("/:userID/personal-details", h.UserHandler.CreatePersonalDetailsHandler)
	rg.PUT("/:userID/personal-details", h.UserHandler.UpdatePersonalDetailsHandler)
	rg.PATCH("/:userID/upload-profile-image", h.UserHandler.UploadProfileImageHandler)
	rg.GET("/search", h.UserHandler.SearchUserHandler)
}
