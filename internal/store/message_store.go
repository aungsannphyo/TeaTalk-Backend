package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type messageRepo struct {
	db *sql.DB
}

func (r *messageRepo) CreateMessage(m *models.Message) error {
	query := `INSERT INTO messages (conversation_id, sender_id, content)
	VALUES (?, ?, ?)
	`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(m.ConversationID, m.SenderID, m.Content)

	if err != nil {
		return err
	}
	return nil
}
