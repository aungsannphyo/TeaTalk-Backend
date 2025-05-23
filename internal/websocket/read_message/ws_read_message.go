package readmessage

import "time"

type WSReadMessage struct {
	MessageID []byte    `json:"messageId"`
	ReaderId  string    `json:"readerId"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
