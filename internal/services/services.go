package services

import (
	s "github.com/aungsannphyo/ywartalk/internal/domain/service"
	store "github.com/aungsannphyo/ywartalk/internal/store/repo"
)

type ServiceFactory interface {
	UserService() s.UserService
	ConversationService() s.ConversationService
	FriendRequestService() s.FriendRequestService
	FriendService() s.FriendService
	MessageService() s.MessageService
	ConversationKeyService() s.ConversationKeyService
}

type serviceFactory struct {
	repoFactory     store.RepositoryFactory
	msgSvc          s.MessageService
	convKeySvc      s.ConversationKeyService
	sessionKeyCache *SessionKeyCache
	userSvc         s.UserService
}

func NewServiceFactory(repoFactory store.RepositoryFactory) ServiceFactory {
	// manually wire after creation

	sessionKeyCache := NewSessionKeyCache()

	convKeySvc := &cKeyService{
		cKeyRepo:        repoFactory.NewConversationKeyRepo(),
		userRepo:        repoFactory.NewUserRepo(),
		sessionKeyCache: sessionKeyCache,
	}

	userSvc := &userServices{
		userRepo:        repoFactory.NewUserRepo(),
		sessionKeyCache: sessionKeyCache,
	}

	messageSvc := &messageService{
		mRepo:      repoFactory.NewMessageRepo(),
		fRepo:      repoFactory.NewFriendRepo(),
		cRepo:      repoFactory.NewConversationRepo(),
		cmRepo:     repoFactory.NewConversationMemberRepo(),
		cKeyRepo:   repoFactory.NewConversationKeyRepo(),
		convKeySvc: convKeySvc,
	}

	return &serviceFactory{
		repoFactory:     repoFactory,
		msgSvc:          messageSvc,
		convKeySvc:      convKeySvc,
		sessionKeyCache: sessionKeyCache,
		userSvc:         userSvc,
	}

}

func (f *serviceFactory) UserService() s.UserService {
	return f.userSvc
}

func (f *serviceFactory) ConversationService() s.ConversationService {
	return &conService{
		cRepo:  f.repoFactory.NewConversationRepo(),
		cmRepo: f.repoFactory.NewConversationMemberRepo(),
		gaRepo: f.repoFactory.NewGroupAdminRepo(),
		giRepo: f.repoFactory.NewGroupInviteRepo(),
		fRepo:  f.repoFactory.NewFriendRepo(),
	}
}

func (f *serviceFactory) FriendRequestService() s.FriendRequestService {
	return &frService{
		frRepo:  f.repoFactory.NewFriendRequestRepo(),
		fRepo:   f.repoFactory.NewFriendRepo(),
		frlRepo: f.repoFactory.NewFriendRequestLogRepo(),
	}
}

func (f *serviceFactory) FriendService() s.FriendService {
	return &fService{
		fRepo:   f.repoFactory.NewFriendRepo(),
		frRepo:  f.repoFactory.NewFriendRequestRepo(),
		frlRepo: f.repoFactory.NewFriendRequestLogRepo(),
	}
}

func (f *serviceFactory) MessageService() s.MessageService {
	return f.msgSvc
}

func (f *serviceFactory) ConversationKeyService() s.ConversationKeyService {
	return f.convKeySvc
}
