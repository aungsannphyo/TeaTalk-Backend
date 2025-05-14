package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type cmRepo struct {
	db *sql.DB
}

func NewConversationMemberRepo(db *sql.DB) *cmRepo {
	return &cmRepo{
		db: db,
	}
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
