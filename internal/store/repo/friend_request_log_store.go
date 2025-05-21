package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type frlRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *frlRepo) CreateFriendRequestLog(frl *models.FriendRequestLog) error {
	query, err := r.loader.LoadQuery("sql/friend_request_log/create_friend_request_log.sql")

	if err != nil {
		return err
	}
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
