package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.GET("/", h.UserHandler.GetUserHandler)
	rg.PUT("/username", h.UserHandler.UpdateUserName)
	rg.GET("/chat-list", h.UserHandler.GetChatListByUserIdHandler)
	rg.PUT("/personal-details", h.UserHandler.UpdatePersonalDetailsHandler)
	rg.PATCH("/upload-profile-image", h.UserHandler.UploadProfileImageHandler)
	rg.GET("/search", h.UserHandler.SearchUserHandler)
	rg.GET("/friend", h.UserHandler.GetFriendsByUserHandler)
}
