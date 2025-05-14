package repository

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type GroupInviteRepository interface {
	CreateGroupInvite(cgi *models.GroupInvite) error
	ModerateGroupInvite(mgi *models.GroupInvite) error
}
