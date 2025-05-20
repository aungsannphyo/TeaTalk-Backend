package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type userRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *userRepo) Register(u *models.User) error {

	query, err := r.loader.LoadQuery("sql/auth/register.sql")

	if err != nil {
		return err
	}

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
	query, err := r.loader.LoadQuery("sql/auth/login.sql")

	if err != nil {
		return nil, err
	}

	row := db.DBInstance.QueryRow(query, user.Email)

	var foundUser models.User

	err = row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Email, &foundUser.Password)

	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func (r *userRepo) GetUserById(ctx context.Context, userID string) (*models.User, error) {
	query, err := r.loader.LoadQuery("sql/user/get_user_by_id.sql")

	if err != nil {
		return nil, err
	}

	row := db.DBInstance.QueryRowContext(ctx, query, userID)

	var user models.User

	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetChatListByUserId(ctx context.Context, userID string) ([]models.ChatListItem, error) {
	query, err := r.loader.LoadQuery("sql/user/get_chat_list_by_id.sql")

	if err != nil {
		return nil, err
	}

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

func (r *userRepo) CreatePersonalDetail(ps *models.PersonalDetails) error {
	query, err := r.loader.LoadQuery("sql/user/create_personal_details.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ps.UserID, ps.Gender, ps.DateOfBirth, ps.Bio)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) UpdatePersonalDetail(ps *models.PersonalDetails) error {
	query, err := r.loader.LoadQuery("sql/user/update_personal_details.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ps.Gender, ps.DateOfBirth, ps.Bio, ps.UserID)

	if err != nil {
		return err
	}

	return nil

}

func (r *userRepo) GetProfileImagePath(ctx context.Context, userID string) (string, error) {
	query, err := r.loader.LoadQuery("sql/user/get_profile_image_by_id.sql")

	if err != nil {
		return "", err
	}

	row := db.DBInstance.QueryRowContext(ctx, query, userID)

	var profileImage string

	err = row.Scan(&profileImage)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return profileImage, nil
}

func (r *userRepo) UploadProfileImage(userID string, imagePath string) error {
	query, err := r.loader.LoadQuery("sql/user/update_profile_image.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(imagePath, userID)

	if err != nil {
		return err
	}

	return nil
}
