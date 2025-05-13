package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Register(u *models.User) error {

	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Username, u.Email, u.Password)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Login(user *models.User) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = ?"

	row := db.DBInstance.QueryRow(query, user.Email)

	var foundUser models.User

	err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Email, &foundUser.Password)

	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func (r *userRepo) GetUser(userId string) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = ?"

	row := db.DBInstance.QueryRow(query, userId)

	var user models.User

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
