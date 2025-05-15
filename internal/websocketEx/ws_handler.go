package websocketEx

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")

		if userID == "" {
			log.Println("Missing user_id in query params")
			c.JSON(http.StatusBadRequest, gin.H{"error": "user_id required"})
			return
		}

		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // For development only. Restrict in production.
			},
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WebSocket upgrade failed: %v", err)
			return
		}

		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), id: userID}
		hub.register <- client

		go client.writePump()
		go client.readPump()
	}
}
