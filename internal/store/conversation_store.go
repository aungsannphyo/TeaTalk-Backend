package store

import (
	"context"
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type conRepo struct {
	db *sql.DB
}

func (r *conRepo) CreateConversation(c *models.Conversation) error {
	query := `INSERT INTO 
	conversations (id,is_group, name, created_by)
	VALUES (? ,?, ?, ?)
	`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(c.ID, c.IsGroup, c.Name, c.CreatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *conRepo) CheckExistsConversation(ctx context.Context, senderId, receiverId string) (models.Conversation, error) {
	query := ` SELECT c.id, c.is_group, c.name, c.created_by, c.created_at
	FROM conversations c
	JOIN conversation_members m1 ON c.id = m1.conversation_id 
	JOIN conversation_members m2 ON c.id = m2.conversation_id 
	WHERE c.is_group = FALSE
	AND m1.user_id = ?
	AND m2.user_id = ?
	AND m1.user_id != m2.user_id
	GROUP BY c.id
	HAVING COUNT(DISTINCT m1.user_id) = 1 AND COUNT(DISTINCT m2.user_id) = 1;
	`

	rows := db.DBInstance.QueryRowContext(ctx, query, senderId, receiverId)

	var c models.Conversation

	err := rows.Scan(&c.ID, &c.IsGroup, &c.Name, &c.CreatedBy, &c.CreatedAt)

	if err != nil {
		return models.Conversation{}, err
	}

	return c, nil
}

func (r *conRepo) UpdateGroupName(c *models.Conversation) error {
	query := `UPDATE conversations 
	SET name = ? 
	WHERE id = ?
	`
	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(c.Name, c.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *conRepo) CheckExistsGroup(ctx context.Context, c *models.Conversation) bool {
	query := `SELECT COUNT(*)
	FROM conversations 
	WHERE id = ? AND is_group = 1
	`

	row := db.DBInstance.QueryRowContext(ctx, query, c.ID)

	var con int64

	err := row.Scan(&con)
	if err != nil {
		return false
	}

	return con > 0
}

func (r *conRepo) GetGroupMembers(ctx context.Context, conversationId string) ([]models.User, error) {
	query := `
	SELECT  u.id, u.username, u.email, u.created_at
	FROM conversation_members cm
	JOIN users u ON cm.user_id = u.id
	JOIN conversations c ON cm.conversation_id = c.id
	WHERE c.id = ? AND c.is_group = TRUE
	`

	rows, err := r.db.QueryContext(ctx, query, conversationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
