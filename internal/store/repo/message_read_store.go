package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type mrRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *mrRepo) CreateReadMessage(mr *models.MessageRead) error {
	query, err := r.loader.LoadQuery("sql/message_read/create_message_read.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {

		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(mr.MessageId, mr.UserID)

	if err != nil {
		return err
	}
	return nil
}

func (r *mrRepo) MarkAllReadMessages(userID, conversationID string) error {
	query, err := r.loader.LoadQuery("sql/message_read/mark_all_read_messages.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {

		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID, userID, conversationID)

	if err != nil {
		return err
	}
	return nil

}
