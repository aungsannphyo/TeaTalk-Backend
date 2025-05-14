package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type cRepo struct {
	db *sql.DB
}

func NewConversationRepo(db *sql.DB) *cRepo {
	return &cRepo{
		db: db,
	}
}

func (r *cRepo) CreateConversation(c *models.Conversation) error {
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

func (r *cRepo) CheckExistsConversation(senderId, receiverId string) ([]models.Conversation, error) {
	query := ` SELECT * FROM conversations c
	JOIN conversation_members m1 ON c.id = m1.conversation_id 
	JOIN conversation_members m2 ON c.id = m2.conversation_id 
	WHERE c.is_group = FALSE
	AND m1.user_id = ?
	AND m2.user_id = ?
	AND m1.user_id != m2.user_id
	GROUP BY c.id
	HAVING COUNT(DISTINCT m1.user_id) = 1 AND COUNT(DISTINCT m2.user_id) = 1;
	`

	rows, err := db.DBInstance.Query(query, senderId, receiverId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var conversations []models.Conversation

	for rows.Next() {
		var c models.Conversation

		err := rows.Scan(&c.ID, &c.IsGroup, &c.Name, &c.CreatedBy, &c.CreatedAt)

		if err != nil {
			return nil, err
		}

		conversations = append(conversations, c)
	}

	return conversations, nil
}
