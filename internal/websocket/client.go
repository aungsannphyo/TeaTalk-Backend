package websocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	UserID string
}

type PrivateMessage struct {
	FromUserID string
	ToUserID   string
	Content    []byte
}

type GroupMessage struct {
	FromUserID string
	GroupID    string
	Content    []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(512)
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
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
			log.Println("Invalid message format:", err)
			continue
		}
		wsMsg.CreatedAt = time.Now().Unix()

		switch wsMsg.Type {
		case "private":
			if wsMsg.ToUserID == "" {
				log.Println("Missing to_user_id for private message")
				continue
			}
			response, _ := json.Marshal(wsMsg)
			c.Hub.privateMessage <- PrivateMessage{
				FromUserID: c.UserID,
				ToUserID:   wsMsg.ToUserID,
				Content:    response,
			}
		case "group":
			if wsMsg.GroupID == "" {
				log.Println("Missing group_id for group message")
				continue
			}
			response, _ := json.Marshal(wsMsg)
			c.Hub.groupMessage <- GroupMessage{
				FromUserID: c.UserID,
				GroupID:    wsMsg.GroupID,
				Content:    response,
			}
		default:
			log.Println("Unknown message type:", wsMsg.Type)
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
