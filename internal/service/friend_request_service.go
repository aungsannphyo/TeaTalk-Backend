package service

import (
	"errors"

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
	pending := s.friendRequestRepo.HasPendingRequest(sfrDto.SenderId, sfrDto.ReceiverId)

	if pending {
		return errors.New("Friend request already pending!")
	}

	//check already friend
	exist := s.friendRequestRepo.AlreadyFriends(sfrDto.SenderId, sfrDto.ReceiverId)

	if exist {
		return errors.New("Already Friend")
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
		fr, err := s.friendRequestRepo.FindById(dfrDto.FriendRequestId)

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
