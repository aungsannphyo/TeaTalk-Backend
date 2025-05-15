package store

import (
	"context"
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type cmRepo struct {
	db *sql.DB
}

func (r *cmRepo) CreateConversationMember(cm *models.ConversationMember) error {
	query := `INSERT INTO 
	conversation_members (conversation_id, user_id )
	VALUES (? , ?)
	`
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
	query := `SELECT COUNT(*) 
	FROM conversation_members
	WHERE conversation_id = ? AND user_id = ?
	`
	row := db.DBInstance.QueryRowContext(ctx, query, cm.ConversationID, cm.UserID)

	var member int64
	err := row.Scan(&member)

	if err != nil {
		return false
	}

	return member > 0
}
