package store

import (
	"context"
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type cmRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *cmRepo) CreateConversationMember(cm *models.ConversationMember) error {
	query, err := r.loader.LoadQuery("sql/conversation_member/create_conversation_member.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(cm.ConversationID, cm.UserID)

	if err != nil {
		return err
	}
	return nil
}

func (r *cmRepo) CheckConversationMember(ctx context.Context, cm *models.ConversationMember) bool {
	query, err := r.loader.LoadQuery("sql/conversation_member/check_conversation_member.sql")

	if err != nil {
		return false
	}

	row := db.DBInstance.QueryRowContext(ctx, query, cm.ConversationID, cm.UserID)

	var count int64
	err = row.Scan(&count)

	if err != nil {
		return false
	}

	return count > 0
}
