package repository

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type GroupAdminRepository interface {
	CreateGroupAdmin(cga *models.GroupAdmin) error
	IsGroupAdmin(ctx context.Context, cId, userId string) (bool, error)
}
