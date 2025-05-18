package store

import (
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type giRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *giRepo) CreateGroupInvite(cgi *models.GroupInvite) error {
	query, err := r.loader.LoadQuery("sql/group_invite/create_group_invite.sql")

	if err != nil {
		return err
	}

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

func (r *giRepo) ModerateGroupInvite(mgi *models.GroupInvite) error {
	query, err := r.loader.LoadQuery("sql/group_invite/moderate_group_invite.sql")

	if err != nil {
		return err
	}

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
