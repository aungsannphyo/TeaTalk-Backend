package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type friendRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *friendRepo) CreateFriendShip(f *models.Friend) error {
	query, err := r.loader.LoadQuery("sql/friend/create_friend_ship.sql")

	if err != nil {
		return err
	}

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
	query, err := r.loader.LoadQuery("sql/friend/make_un_friend.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(f.UserID, f.FriendID, f.UserID, f.FriendID)

	if err != nil {
		return nil
	}

	return nil
}

func (r *friendRepo) AlreadyFriends(ctx context.Context, senderId, receiverId string) bool {
	query, err := r.loader.LoadQuery("sql/friend/already_friends.sql")

	if err != nil {

		return false
	}

	row := db.DBInstance.QueryRowContext(ctx, query, senderId, receiverId, senderId, receiverId)

	var count int64

	err = row.Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		return false
	}

	return count > 0
}
