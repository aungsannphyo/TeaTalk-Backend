package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
	"github.com/aungsannphyo/ywartalk/pkg/utils"
)

type userRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *userRepo) isUserIdentityUnique(identity string) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query, _ := r.loader.LoadQuery("sql/user/is_user_identity_unique.sql")

	var exists bool

	row := db.DBInstance.QueryRowContext(ctx, query, identity)

	_ = row.Scan(&exists)

	return !exists
}

func (r *userRepo) getUniqueUserIdentity(base string) string {
	normalized := utils.NormalizeNameToUsername(base)
	identity := "@" + normalized
	suffix := 1

	for !r.isUserIdentityUnique(identity) {
		identity = fmt.Sprintf("@%s%d", normalized, suffix)
		suffix++
	}

	return identity
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

	userIdentity := r.getUniqueUserIdentity(u.Username)

	_, err = stmt.Exec(u.Username, userIdentity, strings.ToLower(u.Email), u.Password)

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

func (r *userRepo) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
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

func (r *userRepo) GetChatListByUserID(ctx context.Context, userID string) ([]models.ChatListItem, error) {
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

func (r *userRepo) SearchUser(
	ctx context.Context,
	userID string,
	searchInput string,
) (*response.SearchResultResponse, error) {

	query, err := r.loader.LoadQuery("sql/user/search_user_by_email_or_identity.sql")
	if err != nil {
		return nil, err
	}

	row := db.DBInstance.QueryRowContext(
		ctx, query,
		userID,
		userID,
		strings.ToLower(searchInput),
		strings.ToLower(searchInput),
		userID,
	)

	var user response.SearchResultResponse

	err = row.Scan(&user.ID, &user.Email, &user.Username, &user.UserIdentity, &user.IsFriend, &user.ProfileImage)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}
