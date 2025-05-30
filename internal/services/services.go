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
	repoFactory store.RepositoryFactory
	frSvc       s.FriendRequestService
	cSvc        s.ConversationService
}

func NewServiceFactory(repoFactory store.RepositoryFactory) ServiceFactory {
	cSvc := &conService{
		cRepo:  repoFactory.NewConversationRepo(),
		cmRepo: repoFactory.NewConversationMemberRepo(),
		gaRepo: repoFactory.NewGroupAdminRepo(),
		giRepo: repoFactory.NewGroupInviteRepo(),
		fRepo:  repoFactory.NewFriendRepo(),
	}

	frSvc := &frService{
		frRepo:  repoFactory.NewFriendRequestRepo(),
		fRepo:   repoFactory.NewFriendRepo(),
		frlRepo: repoFactory.NewFriendRequestLogRepo(),
		cSvc:    cSvc,
	}

	return &serviceFactory{
		repoFactory: repoFactory,
		cSvc:        cSvc,
		frSvc:       frSvc,
	}

}

func (f *serviceFactory) UserService() s.UserService {
	return &userServices{
		userRepo: f.repoFactory.NewUserRepo(),
	}
}

func (f *serviceFactory) ConversationKeyService() s.ConversationKeyService {
	return &cKeyService{
		cKeyRepo: f.repoFactory.NewConversationKeyRepo(),
	}
}

func (f *serviceFactory) ConversationService() s.ConversationService {
	return f.cSvc
}

func (f *serviceFactory) FriendRequestService() s.FriendRequestService {
	return f.frSvc
}

func (f *serviceFactory) FriendService() s.FriendService {
	return &fService{
		fRepo:   f.repoFactory.NewFriendRepo(),
		frRepo:  f.repoFactory.NewFriendRequestRepo(),
		frlRepo: f.repoFactory.NewFriendRequestLogRepo(),
	}
}

func (f *serviceFactory) MessageService() s.MessageService {
	return &messageService{
		mRepo:    f.repoFactory.NewMessageRepo(),
		fRepo:    f.repoFactory.NewFriendRepo(),
		cRepo:    f.repoFactory.NewConversationRepo(),
		cmRepo:   f.repoFactory.NewConversationMemberRepo(),
		cKeyRepo: f.repoFactory.NewConversationKeyRepo(),
	}
}
