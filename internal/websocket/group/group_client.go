package group

import (
	"context"
	"encoding/json"
	"log"
	"strings"
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
	onlineManager  *ws.SharedOnlineManager
}

type GroupMessage struct {
	SenderID       string
	ConversationID string
	Content        []byte
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

		if len(message) == 0 || len(trimSpaces(string(message))) == 0 {
			continue
		}

		var wsMsg WSGroupMessage
		if err := json.Unmarshal(message, &wsMsg); err != nil {
			log.Printf("Failed to unmarshal WSGroupMessage. Raw message: %s, error: %v", string(message), err)
			continue
		}
		wsMsg.CreatedAt = time.Now()

		if wsMsg.ConversationID == "" {
			log.Printf("Missing conversation id  for group message. Raw message: %s", string(message))
			continue
		}

		if len(wsMsg.Content) == 0 || len(trimSpaces(wsMsg.Content)) == 0 {
			continue
		}

		response, _ := json.Marshal(wsMsg)

		dto := dto.SendGroupMessageDto{
			Content: wsMsg.Content,
		}
		ctx := context.Background()

		if err := c.messageService.SendGroupMessage(ctx, wsMsg.ConversationID, c.userID, dto); err != nil {
			log.Println(err)
		}

		c.hub.broadcast <- GroupMessage{
			SenderID:       c.userID,
			ConversationID: wsMsg.ConversationID,
			Content:        response,
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

func trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
