package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type friendRequestRepo struct {
	db *sql.DB
}

func NewFriendRequestRepo(db *sql.DB) *friendRequestRepo {
	return &friendRequestRepo{
		db: db,
	}
}

func (r *friendRequestRepo) SendFriendRequest(fr *models.FriendRequest) error {
	query := `INSERT INTO 
	friend_requests (sender_id, receiver_id, status) 
	VALUES (?, ?, ?)`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(fr.SenderId, fr.ReceiverId, models.StatusPending)

	if err != nil {
		return err
	}

	return nil

}

func (r *friendRequestRepo) DecideFriendRequest(dfr *models.FriendRequest) error {
	query := `UPDATE friend_requests 
	SET  status = ? 
	WHERE receiver_id = ? AND id = ?`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(dfr.Status, dfr.ReceiverId, dfr.ID)

	return err
}

func (r *friendRequestRepo) FindById(friendRequestId string) (*models.FriendRequest, error) {
	query := "SELECT * FROM friend_requests WHERE id = ?"

	row := db.DBInstance.QueryRow(query, friendRequestId)

	var fr models.FriendRequest

	err := row.Scan(&fr.ID, &fr.SenderId, &fr.ReceiverId, &fr.Status, &fr.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &fr, nil
}

func (r *friendRequestRepo) AlreadyFriends(senderId, receiverId string) (bool, error) {
	query := `SELECT 1 FROM friends 
	WHERE user_id =  ? AND friend_id = ?`

	row := db.DBInstance.QueryRow(query, senderId, receiverId)

	var friend int64

	err := row.Scan(&friend)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r *friendRequestRepo) HasPendingRequest(senderId, receiverId string) (bool, error) {
	query := `SELECT 1 FROM friend_requests 
	WHERE sender_id = ? AND receiver_id  = ? AND status = "PENDING"`

	row := db.DBInstance.QueryRow(query, senderId, receiverId)

	var pending int64

	err := row.Scan(&pending)
	if err != nil {
		return false, nil
	}
	return true, nil
}
