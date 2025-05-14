package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type GroupAdminRepository interface {
	CreateGroupAdmin(cga *models.GroupAdmin) error
	IsGroupAdmin(cId, userId string) (bool, error)
}
