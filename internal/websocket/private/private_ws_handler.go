package private

import (
	"net/http"

	s "github.com/aungsannphyo/ywartalk/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketPrivateHandler struct {
	hub        *PrivateHub
	msgService s.MessageService
}

func NewWebSocketPrivateHandler(
	hub *PrivateHub,
	msgS s.MessageService,
) *WebSocketPrivateHandler {
	return &WebSocketPrivateHandler{
		hub:        hub,
		msgService: msgS,
	}
}

func (h *WebSocketPrivateHandler) WebSocketPrivateHandler(c *gin.Context) {
	userID := c.GetString("userId")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &PrivateClient{
		hub:            h.hub,
		conn:           conn,
		send:           make(chan []byte, 512),
		userID:         userID,
		messageService: h.msgService,
	}

	client.hub.register <- client

	go client.ReadPrivatePump()
	go client.WritePrivatePump()

}
