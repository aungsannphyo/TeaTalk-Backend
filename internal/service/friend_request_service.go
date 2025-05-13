package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	"github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type FriendRequestService struct {
	friendRequestRepo repository.FriendRequestRepository
	friendRepo        repository.FriendRepository
}

func NewFriendRequestService(
	frRepo repository.FriendRequestRepository,
	fRepo repository.FriendRepository,
) *FriendRequestService {
	return &FriendRequestService{
		friendRequestRepo: frRepo,
		friendRepo:        fRepo,
	}
}

func (s *FriendRequestService) SendFriendRequest(sfrDto *dto.SendFriendRequestDto) error {

	//check already send friend request
	pending, err := s.HasPendingRequest(sfrDto.SenderId, sfrDto.ReceiverId)

	if err != nil {
		return err
	}

	//check already friend
	exist, err := s.AlreadyFriends(sfrDto.SenderId, sfrDto.ReceiverId)

	if err != nil {
		return err
	}

	if !pending && !exist {
		//make friend request

		friendRequest := &models.FriendRequest{
			SenderId:   sfrDto.SenderId,
			ReceiverId: sfrDto.ReceiverId,
		}
		return s.friendRequestRepo.SendFriendRequest(friendRequest)
	}

	return nil
}

func (s *FriendRequestService) DecideFriendRequest(dfrDto *dto.DecideFriendRequestDto) error {
	decideFriendRequest := &models.FriendRequest{
		ID:         dfrDto.FriendRequestId,
		ReceiverId: dfrDto.UserId,
		Status:     dfrDto.Status,
	}

	//check decide status is ACCEPTED
	//write into friends database for both friendship
	if models.StatusAccepted == dfrDto.Status {
		fr, err := s.FindById(dfrDto.FriendRequestId)

		if err != nil {
			return err
		}

		f := &models.Friend{
			UserID:   fr.ReceiverId,
			FriendID: fr.SenderId,
		}

		s.friendRepo.CreateFriendShip(f)
	}

	return s.friendRequestRepo.DecideFriendRequest(decideFriendRequest)
}

func (s *FriendRequestService) FindById(frId string) (*models.FriendRequest, error) {
	fr, err := s.friendRequestRepo.FindById(frId)

	if err != nil {
		return nil, err
	}

	return fr, nil
}

func (s *FriendRequestService) AlreadyFriends(senderId, receiverId string) (bool, error) {
	friend, err := s.friendRequestRepo.AlreadyFriends(senderId, receiverId)

	if err != nil {
		return false, nil
	}

	return friend, nil
}

func (s *FriendRequestService) HasPendingRequest(senderId, receiverId string) (bool, error) {
	pending, err := s.friendRequestRepo.HasPendingRequest(senderId, receiverId)

	if err != nil {
		return false, nil
	}

	return pending, nil
}
