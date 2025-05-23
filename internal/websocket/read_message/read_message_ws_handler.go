package readmessage

import (
	"net/http"

	"github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketReadMessageHandler struct {
	hub            *ReadMessageHub
	msgReadService service.MessageReadService
}

func NewWebSocketReadMessageHandler(
	hub *ReadMessageHub,
	msgRead service.MessageReadService,

) *WebSocketReadMessageHandler {
	return &WebSocketReadMessageHandler{
		hub:            hub,
		msgReadService: msgRead,
	}
}

func (h *WebSocketReadMessageHandler) WebSocketReadMessageHandler(c *gin.Context) {
	userID := c.GetString("userID")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &ReadMessageClient{
		hub:            h.hub,
		conn:           conn,
		send:           make(chan []byte, 512),
		userID:         userID,
		msgReadService: h.msgReadService,
	}

	client.hub.register <- client

	go client.ReadMessageReadPump()
	go client.WriteMessageReadPump()

}
