package store

import (
	"context"
	"database/sql"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
	"github.com/aungsannphyo/ywartalk/pkg/db"
)

type gaRepo struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func (r *gaRepo) CreateGroupAdmin(cga *models.GroupAdmin) error {
	query, err := r.loader.LoadQuery("sql/group_admin/create_group_admin.sql")

	if err != nil {
		return err
	}

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(cga.ConversationID, cga.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (r *gaRepo) IsGroupAdmin(ctx context.Context, cId, userId string) (bool, error) {
	query, err := r.loader.LoadQuery("sql/group_admin/is_group_admin.sql")

	if err != nil {
		return false, err
	}

	row := db.DBInstance.QueryRowContext(ctx, query, cId, userId)

	var groupAdmin int64

	if err := row.Scan(&groupAdmin); err != nil {
		return false, err
	}

	return groupAdmin > 0, nil
}
