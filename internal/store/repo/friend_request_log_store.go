package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

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

func (r *frlRepo) HasRejectedFriendRequestLog(frl *models.FriendRequestLog) (bool, error) {
	query, err := r.loader.LoadQuery("sql/friend_request_log/has_reject_friend_request_log.sql")

	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	row := db.DBInstance.QueryRowContext(ctx, query, frl.SenderID, frl.ReceiverID, frl.ReceiverID, frl.SenderID)

	var lastAction string

	err = row.Scan(&lastAction)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, nil
	}

	if lastAction == string(models.ActionRejected) {
		return true, nil
	}

	return false, nil
}
