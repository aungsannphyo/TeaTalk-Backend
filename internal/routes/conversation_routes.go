package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterConversationRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.PUT("/update-conversation/:conversationID", h.ConversationsHandler.UpdateGroupNameHandler)
	rg.POST("/:conversationID/invite", h.ConversationsHandler.InviteGroupHandler)
	rg.PATCH("/:conversationID/invite/:inviteUserID", h.ConversationsHandler.ModerateGroupInviteHandler)
	rg.POST("/:conversationID/assign-admin", h.ConversationsHandler.AssignAdminHandler)
	rg.GET("/member/:conversationID", h.ConversationsHandler.GetGroupMembersHandler)
	rg.GET("/groups", h.ConversationsHandler.GetGroupsByIdHandler)
	rg.GET("/sender/:senderId/receiver/:receiverId", h.ConversationsHandler.GetConversationHandler)
}
