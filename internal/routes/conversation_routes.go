package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterConversationRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.POST("/create-group", h.ConversationsHandler.CreateGroupHandler)
	rg.PUT("/update-group/:groupId", h.ConversationsHandler.UpdateGroupNameHandler)
	rg.POST("/:groupId/invite", h.ConversationsHandler.InviteGroupHandler)
	rg.PATCH("/:groupId/invite/:inviteUserId", h.ConversationsHandler.ModerateGroupInviteHandler)
	rg.POST("/:groupId/assign-admin", h.ConversationsHandler.AssignAdminHandler)
	rg.GET("/member/:groupId", h.ConversationsHandler.GetGroupMembersHandler)
}
