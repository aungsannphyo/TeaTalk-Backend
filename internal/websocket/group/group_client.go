package group

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	ws "github.com/aungsannphyo/ywartalk/internal/websocket"
	"github.com/gorilla/websocket"
)

type GroupClient struct {
	hub            *GroupHub
	conn           *websocket.Conn
	send           chan []byte
	userID         string
	messageService service.MessageService
}

type GroupMessage struct {
	SenderID string
	GroupID  string
	Content  []byte
}

func (c *GroupClient) ReadGroupPump() {
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

		var wsMsg WSGroupMessage
		if err := json.Unmarshal(message, &wsMsg); err != nil {
			log.Println("Failed to marshal WSMessage:", err)
			continue
		}
		wsMsg.CreatedAt = time.Now()

		if wsMsg.GroupID == "" {
			log.Println("Missing group_id for group message")
			continue
		}

		response, _ := json.Marshal(wsMsg)

		dto := dto.SendGroupMessageDto{
			Content: wsMsg.Content,
		}
		ctx := context.Background()

		if err := c.messageService.SendGroupMessage(ctx, wsMsg.GroupID, c.userID, dto); err != nil {
			log.Println(err)
		}

		c.hub.broadcast <- GroupMessage{
			SenderID: c.userID,
			GroupID:  wsMsg.GroupID,
			Content:  response,
		}
	}
}

func (c *GroupClient) WriteGroupPump() {
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
