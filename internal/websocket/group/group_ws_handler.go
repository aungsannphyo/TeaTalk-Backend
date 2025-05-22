package group

import (
	"net/http"

	s "github.com/aungsannphyo/ywartalk/internal/domain/service"

	ws "github.com/aungsannphyo/ywartalk/internal/websocket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketGroupHandler struct {
	hub           *GroupHub
	msgService    s.MessageService
	onlineManager *ws.SharedOnlineManager
}

func NewWebSocketGroupHandler(
	hub *GroupHub,
	msgS s.MessageService,
	onlineManager *ws.SharedOnlineManager,
) *WebSocketGroupHandler {
	return &WebSocketGroupHandler{
		hub:           hub,
		msgService:    msgS,
		onlineManager: onlineManager,
	}
}

func (h *WebSocketGroupHandler) NewWebSocketGroupHandler(c *gin.Context) {
	userID := c.GetString("userID")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &GroupClient{
		hub:            h.hub,
		conn:           conn,
		send:           make(chan []byte, 512),
		userID:         userID,
		messageService: h.msgService,
		onlineManager:  h.onlineManager,
	}

	client.hub.register <- client

	go client.ReadGroupPump()
	go client.WriteGroupPump()

}
