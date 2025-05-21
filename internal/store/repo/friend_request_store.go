package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type frRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *frRepo) SendFriendRequest(fr *models.FriendRequest) error {
	query, err := r.loader.LoadQuery("sql/friend_request/send_friend_request.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(fr.SenderId, fr.ReceiverId, models.FriendRequestPending)

	if err != nil {
		return err
	}

	return nil
}

func (r *frRepo) RejectFriendRequest(dfr *models.FriendRequest) error {
	query, err := r.loader.LoadQuery("sql/friend_request/reject_friend_request.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(dfr.Status, dfr.ReceiverId, dfr.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *frRepo) GetFriendRequestByID(ctx context.Context, id string) (*models.FriendRequest, error) {
	query, err := r.loader.LoadQuery("sql/friend_request/get_friend_request_by_id.sql")

	if err != nil {
		return nil, err
	}

	row := db.DBInstance.QueryRowContext(ctx, query, id)

	var fr models.FriendRequest

	err = row.Scan(&fr.ID, &fr.SenderId, &fr.ReceiverId, &fr.Status, &fr.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &fr, nil
}

func (r *frRepo) DeleteFriendRequestByID(id string) error {
	query, err := r.loader.LoadQuery("sql/friend_request/delete_friend_request_by_id.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil
}

func (r *frRepo) HasPendingRequest(ctx context.Context, senderId, receiverId string) bool {

	query, err := r.loader.LoadQuery("sql/friend_request/has_pending_request.sql")

	if err != nil {
		return false
	}

	row := db.DBInstance.QueryRowContext(ctx, query, senderId, receiverId, senderId, receiverId)

	var pending int64

	err = row.Scan(&pending)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		return false
	}
	return pending > 0
}

func (r *frRepo) GetAllFriendRequestLog(ctx context.Context, userID string) (
	[]response.FriendRequestResponse, error,
) {
	query, err := r.loader.LoadQuery("sql/friend_request/get_all_friend_request.sql")

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
