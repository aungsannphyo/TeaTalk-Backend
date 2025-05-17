package websocket

import (
	"net/http"

	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketHandler struct {
	hub        *Hub
	msgService s.MessageService
	uService   s.UserService
}

func NewWebSocketHandler(
	hub *Hub,
	msgS s.MessageService,
	userS s.UserService,
) *WebSocketHandler {
	return &WebSocketHandler{
		hub:        hub,
		msgService: msgS,
		uService:   userS,
	}
}

func (h *WebSocketHandler) WebSocketHandler(c *gin.Context) {
	userID := c.GetString("userId")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &Client{
		Hub:            h.hub,
		Conn:           conn,
		Send:           make(chan []byte, 512),
		UserID:         userID,
		MessageService: h.msgService,
		UserService:    h.uService,
	}

	client.Hub.register <- client

	go client.WritePump()
	go client.ReadPump()

}
