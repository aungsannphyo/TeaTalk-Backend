package private

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

type PrivateClient struct {
	hub            *PrivateHub
	conn           *websocket.Conn
	send           chan []byte
	userID         string
	messageService service.MessageService
}

type PrivateMessage struct {
	SenderID   string
	ReceiverID string
	Content    []byte
}

func (c *PrivateClient) ReadPrivatePump() {
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

		var wsMsg WSPrivateMessage
		if err := json.Unmarshal(message, &wsMsg); err != nil {
			log.Println("Failed to marshal WSMessage:", err)
			continue
		}
		wsMsg.CreatedAt = time.Now()

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
		if err := c.messageService.SendPrivateMessage(ctx, c.userID, dto); err != nil {
			log.Println(err)
		}

		c.hub.broadcast <- PrivateMessage{
			SenderID:   c.userID,
			ReceiverID: wsMsg.ReceiverID,
			Content:    response,
		}
	}
}

func (c *PrivateClient) WritePrivatePump() {
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
