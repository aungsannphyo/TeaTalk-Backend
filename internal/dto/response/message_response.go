package response

import (
	"database/sql"
	"time"
)

type MessageResponse struct {
	ConversationID   string
	SenderID         string
	ReceiverID       string
	Content          string
	IsRead           bool
	SeenByName       sql.NullString
	MessageCreatedAt time.Time
}
