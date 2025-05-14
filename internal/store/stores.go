package store

import (
	"database/sql"

	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
)

type RepositoryFactory interface {
	NewUserRepo() r.UserRepository
	NewFriendRequestRepo() r.FriendRequestRepository
	NewFriendRepo() r.FriendRepository
	NewFriendRequestLogRepo() r.FriendRequestLogRepository
	NewMessageRepo() r.MessageRepository
	NewConversationRepo() r.ConversationRepository
	NewConversationMemberRepo() r.ConversationMemeberRepository
	NewGroupAdminRepo() r.GroupAdminRepository
	NewGroupInviteRepo() r.GroupInviteRepository
}

type factory struct {
	db *sql.DB
}

func NewRepositoryFactory(db *sql.DB) RepositoryFactory {
	return &factory{db: db}
}

func (f *factory) NewUserRepo() r.UserRepository {
	return &userRepo{db: f.db}
}

func (f *factory) NewFriendRequestRepo() r.FriendRequestRepository {
	return &frRepo{db: f.db}
}

func (f *factory) NewFriendRepo() r.FriendRepository {
	return &friendRepo{db: f.db}
}

func (f *factory) NewFriendRequestLogRepo() r.FriendRequestLogRepository {
	return &frlRepo{db: f.db}
}

func (f *factory) NewMessageRepo() r.MessageRepository {
	return &messageRepo{db: f.db}
}

func (f *factory) NewConversationRepo() r.ConversationRepository {
	return &conRepo{db: f.db}
}

func (f *factory) NewConversationMemberRepo() r.ConversationMemeberRepository {
	return &cmRepo{db: f.db}
}

func (f *factory) NewGroupAdminRepo() r.GroupAdminRepository {
	return &gaRepo{db: f.db}
}

func (f *factory) NewGroupInviteRepo() r.GroupInviteRepository {
	return &giRepo{db: f.db}
}
