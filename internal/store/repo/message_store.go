package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type messageRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *messageRepo) CreateMessage(m *models.Message) error {
	query, err := r.loader.LoadQuery("sql/message/create_message.sql")

	if err != nil {
		return err
	}

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
