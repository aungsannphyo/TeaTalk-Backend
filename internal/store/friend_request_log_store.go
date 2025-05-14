package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type frlRepo struct {
	db *sql.DB
}

func (r *frlRepo) CreateFriendRequestLog(frl *models.FriendRequestLog) error {
	query := `INSERT INTO 
	friend_request_logs (sender_id, receiver_id, action, performed_by)
	VALUES (?, ?, ?, ?)
	`
	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(frl.SenderID, frl.ReceiverID, frl.Action, frl.PerformedBy)

	if err != nil {
		return err
	}

	return nil
}
