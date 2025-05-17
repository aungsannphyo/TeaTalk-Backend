package store

import (
	"context"
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type userRepo struct {
	db *sql.DB
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

func (r *userRepo) GetUserById(ctx context.Context, userID string) (*models.User, error) {
	query := "SELECT id, username, email, password, created_at FROM users WHERE id = ?"

	row := db.DBInstance.QueryRowContext(ctx, query, userID)

	var user models.User

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetGroupsById(ctx context.Context, userID string) ([]models.Conversation, error) {
	query := `
		SELECT c.id, c.is_group, c.name, c.created_by, c.created_at
		FROM conversation_members cm
		JOIN conversations c ON cm.conversation_id = c.id
		WHERE cm.user_id = ?
		AND c.is_group = TRUE 
	`

	rows, err := db.DBInstance.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conversations []models.Conversation
	for rows.Next() {
		var c models.Conversation
		if err := rows.Scan(&c.ID, &c.IsGroup, &c.Name, &c.CreatedBy, &c.CreatedAt); err != nil {
			return nil, err
		}
		conversations = append(conversations, c)
	}

	return conversations, nil
}
