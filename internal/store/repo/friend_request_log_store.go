package store

import (
	"context"
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type frlRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *frlRepo) CreateFriendRequestLog(frl *models.FriendRequestLog) error {
	query, err := r.loader.LoadQuery("sql/friend_request_log/create_friend_request.sql")

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

func (r *frlRepo) GetAllFriendRequestLog(ctx context.Context, userID string) (
	[]response.FriendRequestResponse, error,
) {
	query, err := r.loader.LoadQuery("sql/friend_request_log/get_all_friend_request_log.sql")

	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []response.FriendRequestResponse

	for rows.Next() {
		var log response.FriendRequestResponse
		if err := rows.Scan(
			&log.RequestID,
			&log.SenderID,
			&log.Username,
			&log.Email,
			&log.ProfileImage,
			&log.CreateAt,
		); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil

}
