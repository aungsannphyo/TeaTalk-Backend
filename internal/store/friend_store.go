package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type friendRepo struct {
	db *sql.DB
}

func (r *friendRepo) CreateFriendShip(f *models.Friend) error {
	query := `INSERT INTO friends (user_id, friend_id ) VALUES (?, ?)`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(f.UserID, f.FriendID)

	if err != nil {
		return err
	}
	return nil
}

func (r *friendRepo) MakeUnFriend(f *models.Friend) error {
	//need to
	query := "DELETE FROM friends WHERE (user_id = ? AND friend_id = ?) OR (friend_id = ? AND user_id = ?)"

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, firstErr := stmt.Exec(f.UserID, f.FriendID, f.UserID, f.FriendID)
	_, secondErr := stmt.Exec(f.FriendID, f.UserID, f.FriendID, f.UserID)

	if firstErr != nil || secondErr != nil {
		return nil
	}

	return nil
}

func (r *friendRepo) AlreadyFriends(senderId, receiverId string) bool {
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
