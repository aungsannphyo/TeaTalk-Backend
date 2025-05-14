package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type groupInviteRepo struct {
	db *sql.DB
}

func NewGroupInviteRepo(db *sql.DB) *groupInviteRepo {
	return &groupInviteRepo{
		db: db,
	}
}

func (r *groupInviteRepo) CreateGroupInvite(cgi *models.GroupInvite) error {
	query := `INSERT INTO 
	group_invites (conversation_id, invited_by, invited_user_id, status) 
	VALUES (?, ?, ?, ?)
	`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(cgi.ConversationID, cgi.InvitedBy, cgi.InvitedUserId, cgi.Status)

	if err != nil {
		return err
	}

	return nil
}

func (r *groupInviteRepo) ModerateGroupInvite(mgi *models.GroupInvite) error {
	query := "UPDATE group_invites SET status = ? WHERE conversation_id = ? AND invited_user_id = ?"

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(mgi.Status, mgi.ConversationID, mgi.InvitedUserId)

	if err != nil {
		return err
	}

	return nil
}
