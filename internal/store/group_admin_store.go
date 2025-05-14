package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type groupAdminRepo struct {
	db *sql.DB
}

func NewGroupAdminRepo(db *sql.DB) *groupAdminRepo {
	return &groupAdminRepo{
		db: db,
	}
}

func (r *groupAdminRepo) CreateGroupAdmin(cga *models.GroupAdmin) error {
	query := "INSERT INTO group_admins (conversation_id, user_id) VALUES (? , ?)"

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(cga.ConversationID, cga.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (r *groupAdminRepo) IsGroupAdmin(cId, userId string) (bool, error) {
	query := `SELECT COUNT(*) 
	FROM group_admins 
	WHERE conversation_id = ? AND user_id = ? 
	`
	row := db.DBInstance.QueryRow(query, cId, userId)

	var groupAdmin int64

	if err := row.Scan(&groupAdmin); err != nil {
		return false, err
	}

	return groupAdmin > 0, nil
}
