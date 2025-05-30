package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type conRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *conRepo) CreateConversation(c *models.Conversation) error {
	query, err := r.loader.LoadQuery("sql/conversation/create_conversation.sql")

	if err != nil {
		return err
	}

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

func (r *conRepo) GetConversation(ctx context.Context, senderId, receiverId string) (models.Conversation, error) {
	query, err := r.loader.LoadQuery("sql/conversation/check_exists_conversation.sql")

	if err != nil {
		return models.Conversation{}, err
	}
	row := db.DBInstance.QueryRowContext(ctx, query, senderId, receiverId)

	var c models.Conversation
	err = row.Scan(&c.ID, &c.IsGroup, &c.Name, &c.CreatedBy, &c.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Conversation{}, nil
		}
		return models.Conversation{}, err
	}

	return c, nil
}

func (r *conRepo) UpdateGroupName(c *models.Conversation) error {
	query, err := r.loader.LoadQuery("sql/conversation/update_conversation.sql")

	if err != nil {
		return err
	}

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
	query, err := r.loader.LoadQuery("sql/conversation/check_exists_group.sql")

	if err != nil {
		return false
	}

	row := db.DBInstance.QueryRowContext(ctx, query, c.ID)

	var con int64

	err = row.Scan(&con)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		return false
	}

	return con > 0
}

func (r *conRepo) GetGroupMembers(ctx context.Context, conversationId string) ([]models.User, error) {
	query, err := r.loader.LoadQuery("sql/conversation/get_group_members.sql")

	if err != nil {
		return nil, err
	}

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

func (r *conRepo) GetGroupsById(ctx context.Context, userID string) ([]models.Conversation, error) {
	query, err := r.loader.LoadQuery("sql/conversation/get_group_by_id.sql")

	if err != nil {
		return nil, err
	}

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
