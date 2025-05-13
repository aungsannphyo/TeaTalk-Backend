package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type friendRepo struct {
	db *sql.DB
}

func NewFriendRepo(db *sql.DB) *friendRepo {
	return &friendRepo{
		db: db,
	}
}

func (r *friendRepo) CreateFriendShip(f *models.Friend) error {
	query := `INSERT INTO friends (user_id, friend_id ) VALUES (?, ?)`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	//insert two row [bidirectional]
	_, firstErr := stmt.Exec(f.UserID, f.FriendID)
	_, secondErr := stmt.Exec(f.FriendID, f.UserID)

	if firstErr != nil || secondErr != nil {
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
