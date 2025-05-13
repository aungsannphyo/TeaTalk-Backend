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

func (r *friendRequestRepo) RejectFriendRequest(dfr *models.FriendRequest) error {
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

func (r *friendRequestRepo) FindById(id string) (*models.FriendRequest, error) {
	query := "SELECT * FROM friend_requests WHERE id = ?"

	row := db.DBInstance.QueryRow(query, id)

	var fr models.FriendRequest

	err := row.Scan(&fr.ID, &fr.SenderId, &fr.ReceiverId, &fr.Status, &fr.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &fr, nil
}

func (r *friendRequestRepo) DeleteById(id string) error {
	query := "DELETE FROM friend_requests WHERE id = ? "

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

func (r *friendRequestRepo) AlreadyFriends(senderId, receiverId string) bool {
	query := `SELECT COUNT(*) FROM friends 
		WHERE (user_id = ? AND friend_id = ?) OR (friend_id = ? AND user_id = ?)`

	row := db.DBInstance.QueryRow(query, senderId, receiverId, senderId, receiverId)

	var friend int64

	err := row.Scan(&friend)
	if err != nil {
		return false
	}

	return friend > 0
}

func (r *friendRequestRepo) HasPendingRequest(senderId, receiverId string) bool {

	query := `SELECT COUNT(*) FROM friend_requests 
	WHERE ((sender_id = ? AND receiver_id = ?) OR
		   (receiver_id = ? AND sender_id = ?))
	AND status = "PENDING"`

	row := db.DBInstance.QueryRow(query, senderId, receiverId, senderId, receiverId)

	var pending int64

	err := row.Scan(&pending)
	if err != nil {
		return false
	}
	return pending > 0
}
