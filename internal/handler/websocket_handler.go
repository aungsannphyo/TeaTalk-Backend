package handler

import (
	"github.com/aungsannphyo/ywartalk/internal/websocket"
	"github.com/gin-gonic/gin"
)

type WebSocketHandler struct {
	hub *websocket.Hub
}

func NewWebSocketHandler(hub *websocket.Hub) *WebSocketHandler {
	return &WebSocketHandler{
		hub: hub,
	}
}

func (h *WebSocketHandler) WebSocketHandler(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	websocket.HandleWebSocket(c.Writer, c.Request, userID.(string), h.hub)
}
