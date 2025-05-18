package store

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

//go:embed sql/**/*.sql
var sqlFiles embed.FS

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) Register(u *models.User) error {

	queryBytes, err := sqlFiles.ReadFile("sql/auth/register.sql")

	if err != nil {
		log.Fatal(err)
	}

	query := string(queryBytes)
	fmt.Println("SQL query:", query)

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
	queryBytes, err := sqlFiles.ReadFile("sql/auth/login.sql")

	query := string(queryBytes)

	row := db.DBInstance.QueryRow(query, user.Email)

	var foundUser models.User

	err = row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Email, &foundUser.Password)

	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func (r *userRepo) GetUserById(ctx context.Context, userID string) (*models.User, error) {
	queryBytes, err := sqlFiles.ReadFile("sql/user/get_user_by_id.sql")

	query := string(queryBytes)

	row := db.DBInstance.QueryRowContext(ctx, query, userID)

	var user models.User

	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetChatListByUserId(ctx context.Context, userID string) ([]models.ChatListItem, error) {
	queryBytes, err := sqlFiles.ReadFile("sql/user/get_chat_list_by_id.sql")

	query := string(queryBytes)

	rows, err := db.DBInstance.QueryContext(ctx, query, userID, userID, userID, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatList []models.ChatListItem
	for rows.Next() {
		var chat models.ChatListItem
		if err := rows.Scan(
			&chat.ConversationID,
			&chat.IsGroup,
			&chat.Name,
			&chat.LastMessageID,
			&chat.LastMessageContent,
			&chat.LastMessageSender,
			&chat.LastMessageCreatedAt,
			&chat.UnreadCount,
		); err != nil {

			return nil, err
		}
		chatList = append(chatList, chat)
	}

	return chatList, nil

}
