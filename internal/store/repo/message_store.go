package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto"
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

	_, err = stmt.Exec(m.ConversationID, m.SenderID, m.Content, m.MessageNonce)

	if err != nil {
		return err
	}
	return nil
}

func (r *messageRepo) GetMessages(
	ctx context.Context,
	conversationID string,
	cursorTimestamp *time.Time,
	pageSize int,
) ([]dto.MessagesDto, error) {

	query, err := r.loader.LoadQuery("sql/message/get_messages.sql")

	if err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if cursorTimestamp != nil {
		rows, err = db.DBInstance.QueryContext(ctx, query, conversationID, *cursorTimestamp, pageSize)
	} else {
		rows, err = db.DBInstance.QueryContext(ctx, query, conversationID, nil, pageSize)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []dto.MessagesDto{}
	for rows.Next() {
		var m dto.MessagesDto
		err := rows.Scan(
			&m.MessageID,
			&m.MemberID,
			&m.SenderID,
			&m.Content,
			&m.IsRead,
			&m.SeenByName,
			&m.MessageCreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
