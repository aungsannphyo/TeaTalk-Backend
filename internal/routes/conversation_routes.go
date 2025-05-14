package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterConversationRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.POST("/create-group", h.ConversationsHandler.CreateGroup)
	rg.PUT("/update-group/:groupId", h.ConversationsHandler.UpdateGroupName)
	rg.POST("/:groupId/invite", h.ConversationsHandler.InviteGroup)
	rg.PATCH("/:groupId/invite/:inviteUserId", h.ConversationsHandler.ModerateGroupInvite)
	rg.POST("/:groupId/assign-admin", h.ConversationsHandler.AssignAdmin)
}
