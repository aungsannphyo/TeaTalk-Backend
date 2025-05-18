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
}

type serviceFactory struct {
	repoFactory store.RepositoryFactory
}

func NewServiceFactory(repoFactory store.RepositoryFactory) ServiceFactory {
	return &serviceFactory{repoFactory: repoFactory}
}

func (f *serviceFactory) UserService() s.UserService {
	return &userServices{userRepo: f.repoFactory.NewUserRepo()}
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
	return &messageService{
		mRepo:  f.repoFactory.NewMessageRepo(),
		fRepo:  f.repoFactory.NewFriendRepo(),
		cRepo:  f.repoFactory.NewConversationRepo(),
		cmRepo: f.repoFactory.NewConversationMemberRepo(),
	}
}
