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
	hub         *PrivateHub
	msgService  s.MessageService
	userService s.UserService
}

func NewWebSocketPrivateHandler(
	hub *PrivateHub,
	msgS s.MessageService,
	userS s.UserService,
) *WebSocketPrivateHandler {
	return &WebSocketPrivateHandler{
		hub:         hub,
		msgService:  msgS,
		userService: userS,
	}
}

func (h *WebSocketPrivateHandler) WebSocketPrivateHandler(c *gin.Context) {
	userID := c.GetString("userID")

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
		userService:    h.userService,
	}

	client.hub.register <- client

	go client.ReadPrivatePump()
	go client.WritePrivatePump()

}
