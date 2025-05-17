package routes

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterConversationRoutes(rg *gin.RouterGroup, h *handler.HandlerSet) {
	rg.Use(middleware.Middleware)
	rg.POST("/create-group", h.ConversationsHandler.CreateGroupHandler)
	rg.PUT("/update-group/:groupID", h.ConversationsHandler.UpdateGroupNameHandler)
	rg.POST("/:groupID/invite", h.ConversationsHandler.InviteGroupHandler)
	rg.PATCH("/:groupID/invite/:inviteUserID", h.ConversationsHandler.ModerateGroupInviteHandler)
	rg.POST("/:groupID/assign-admin", h.ConversationsHandler.AssignAdminHandler)
	rg.GET("/member/:groupID", h.ConversationsHandler.GetGroupMembersHandler)
	rg.GET("/", h.ConversationsHandler.GetGroupsByIdHandler)
}
