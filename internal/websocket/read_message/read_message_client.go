package readmessage

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/service"
	ws "github.com/aungsannphyo/ywartalk/internal/websocket"
	"github.com/gorilla/websocket"
)

type ReadMessageClient struct {
	hub            *ReadMessageHub
	conn           *websocket.Conn
	send           chan []byte
	userID         string
	msgReadService service.MessageReadService
}

func (c *ReadMessageClient) ReadMessageReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(ws.MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(ws.PongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(ws.PongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		if len(message) == 0 || len(trimSpaces(string(message))) == 0 {
			continue
		}

		var wsMsg WSReadMessage
		if err := json.Unmarshal(message, &wsMsg); err != nil {
			log.Printf("Failed to unmarshal WSPrivateMessage. Raw message: %s, error: %v", string(message), err)
			continue
		}
		wsMsg.CreatedAt = time.Now()

		if wsMsg.ReaderId == "" {
			log.Printf("Missing readerId for read message. Raw message: %s", string(message))
			continue
		}

		if len(wsMsg.ReaderId) == 0 {
			continue
		}

		response, _ := json.Marshal(wsMsg)

		msgRead := &models.MessageRead{
			MessageId: string(wsMsg.MessageID),
			UserID:    wsMsg.ReaderId,
		}

		// insert into database
		if err := c.msgReadService.CreateReadMessage(msgRead); err != nil {
			log.Println(err)
		}

		c.hub.broadcast <- WSReadMessage{
			MessageID: response,
			ReaderId:  wsMsg.ReaderId,
		}
	}
}

func (c *ReadMessageClient) WriteMessageReadPump() {
	ticker := time.NewTicker(ws.PingPeriod)
	defer func() {
		ticker.Stop()
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(ws.WriteWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(message)
			if err != nil {
				return
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(ws.WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
