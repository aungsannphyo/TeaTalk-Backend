package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type cKeyRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *cKeyRepo) CreateConversationKey(cKey *models.ConversationKey) error {
	query, err := r.loader.LoadQuery("sql/conversation_key/create_conversation_key.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(cKey.ConversationId, cKey.UserID, cKey.ConversationEncryptedKey, cKey.ConversationKeyNonce)

	if err != nil {
		return err
	}

	return nil
}

func (r *cKeyRepo) GetConversationKey(ctx context.Context, conversationID, userID string) ([]byte, []byte, error) {
	query, err := r.loader.LoadQuery("sql/conversation_key/get_conversation_key.sql")

	if err != nil {
		return nil, nil, err
	}

	row := db.DBInstance.QueryRowContext(ctx, query, conversationID, userID)

	var encryptedKey []byte
	var nonce []byte

	err = row.Scan(&encryptedKey, &nonce)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, err
		}
		return nil, nil, err
	}

	return encryptedKey, nonce, nil
}
