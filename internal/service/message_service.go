package service

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	r "github.com/aungsannphyo/ywartalk/internal/domain/repository"
	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessageService struct {
	mRepo  r.MessageRepository
	fRepo  r.FriendRepository
	cRepo  r.ConversationRepository
	cmRepo r.ConversationMemeberRepository
}

func NewMessageService(
	mRepo r.MessageRepository,
	fRepo r.FriendRepository,
	cRepo r.ConversationRepository,
	cmRepo r.ConversationMemeberRepository,
) *MessageService {
	return &MessageService{
		mRepo:  mRepo,
		fRepo:  fRepo,
		cRepo:  cRepo,
		cmRepo: cmRepo,
	}
}

func (s *MessageService) SendPrivateMessage(dto dto.SendPrivateMessageDto, c *gin.Context) error {
	senderId := c.GetString("userId")

	//check already friend
	exist := s.fRepo.AlreadyFriends(senderId, dto.ReceiverId)

	//if exists
	if exist {
		cID := uuid.NewString()
		//check already have conversation
		conversations, err := s.cRepo.CheckExistsConversation(senderId, dto.ReceiverId)
		if err != nil {
			return &common.InternalServerError{Message: err.Error()}
		}

		//insert for first time sending message to create conversations
		if len(conversations) != 2 {

			conversation := &models.Conversation{
				ID:        cID,
				IsGroup:   false,
				Name:      nil,
				CreatedBy: nil,
			}
			err := s.cRepo.CreateConversation(conversation)

			if err != nil {
				return &common.InternalServerError{Message: err.Error()}
			}

			//insert into conversation member for both sender and receiver
			cmSender := &models.ConversationMember{
				ConversationID: cID,
				UserID:         senderId,
			}

			cmReceiver := &models.ConversationMember{
				ConversationID: cID,
				UserID:         dto.ReceiverId,
			}

			// [bidirectional]
			senderErr := s.cmRepo.CreateConversationMember(cmSender)
			receiverErr := s.cmRepo.CreateConversationMember(cmReceiver)

			if senderErr != nil || receiverErr != nil {
				return &common.InternalServerError{Message: "Something went wrong, Please try again later"}
			}
		}

		//send the message here if both 2 user are connected as friends
		m := &models.Message{
			ConversationID: cID,
			SenderID:       senderId,
			Content:        dto.Content,
		}

		return s.mRepo.CreateMessage(m)

	} else {
		return &common.UnAuthorizedError{Message: "You cant send the message right now!"}
	}
}
