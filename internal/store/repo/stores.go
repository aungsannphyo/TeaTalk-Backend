package store

import (
	"database/sql"

	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	sqlloader "github.com/aungsannphyo/ywartalk/internal/store/sql_loader"
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
	NewConversationKeyRepo() r.ConversationKeyRepository
}

type repofactory struct {
	db     *sql.DB
	loader sqlloader.SQLLoader
}

func NewRepositoryFactory(db *sql.DB, loader sqlloader.SQLLoader) RepositoryFactory {
	return &repofactory{db: db, loader: loader}
}

func (f *repofactory) NewConversationKeyRepo() r.ConversationKeyRepository {
	return &cKeyRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewUserRepo() r.UserRepository {
	return &userRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewFriendRequestRepo() r.FriendRequestRepository {
	return &frRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewFriendRepo() r.FriendRepository {
	return &friendRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewFriendRequestLogRepo() r.FriendRequestLogRepository {
	return &frlRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewMessageRepo() r.MessageRepository {
	return &messageRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewConversationRepo() r.ConversationRepository {
	return &conRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewConversationMemberRepo() r.ConversationMemeberRepository {
	return &cmRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewGroupAdminRepo() r.GroupAdminRepository {
	return &gaRepo{db: f.db, loader: f.loader}
}

func (f *repofactory) NewGroupInviteRepo() r.GroupInviteRepository {
	return &giRepo{db: f.db, loader: f.loader}
}
