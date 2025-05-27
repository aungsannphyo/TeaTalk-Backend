package service

import (
	"context"
	"time"

	"github.com/aungsannphyo/ywartalk/internal/dto"
	"github.com/aungsannphyo/ywartalk/internal/dto/response"
)

type MessageService interface {
	SendPrivateMessage(ctx context.Context, senderId string, dto dto.SendPrivateMessageDto) error
	SendGroupMessage(ctx context.Context, cID string, userID string, dto dto.SendGroupMessageDto) error
	GetMessages(
		ctx context.Context,
		conversationID string,
		userID string,
		cursorTimestamp *time.Time,
		pageSize int,
	) ([]response.MessageResponse, error)
}
