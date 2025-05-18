package websocket

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/service"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub            *Hub
	Conn           *websocket.Conn
	Send           chan []byte
	UserID         string
	MessageService service.MessageService
	UserService    service.UserService
}

type PrivateMessage struct {
	SenderID   string
	ReceiverID string
	Content    []byte
}

type GroupMessage struct {
	SenderID string
	GroupID  string
	Content  []byte
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var wsMsg WSMessage
		if err := json.Unmarshal(message, &wsMsg); err != nil {
			log.Println("Failed to marshal WSMessage:", err)
			continue
		}
		wsMsg.CreatedAt = time.Now()

		switch wsMsg.Type {
		case "private":
			if wsMsg.ReceiverID == "" {
				log.Println("Missing to_user_id for private message")
				continue
			}

			response, _ := json.Marshal(wsMsg)

			dto := dto.SendPrivateMessageDto{
				ReceiverId: wsMsg.ReceiverID,
				Content:    wsMsg.Content,
			}
			ctx := context.Background()
			// insert into database
			if err := c.MessageService.SendPrivateMessage(ctx, c.UserID, dto); err != nil {
				log.Println(err)
			}

			c.Hub.privateMessage <- PrivateMessage{
				SenderID:   c.UserID,
				ReceiverID: wsMsg.ReceiverID,
				Content:    response,
			}

		case "group":
			if wsMsg.GroupID == "" {
				log.Println("Missing group_id for group message")
				continue
			}

			response, _ := json.Marshal(wsMsg)

			dto := dto.SendGroupMessageDto{
				Content: wsMsg.Content,
			}
			ctx := context.Background()

			if err := c.MessageService.SendGroupMessage(ctx, wsMsg.GroupID, c.UserID, dto); err != nil {
				log.Println(err)
			}

			c.Hub.groupMessage <- GroupMessage{
				SenderID: c.UserID,
				GroupID:  wsMsg.GroupID,
				Content:  response,
			}
		default:
			log.Println("Unknown message type:", wsMsg.Type)
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
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
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
