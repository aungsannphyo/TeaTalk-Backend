package service

import (
	"context"

	"github.com/aungsannphyo/ywartalk/internal/dto"
)

type MessageService interface {
	SendPrivateMessage(ctx context.Context, senderId string, dto dto.SendPrivateMessageDto) error
	SendGroupMessage(ctx context.Context, cID string, userID string, dto dto.SendGroupMessageDto) error
}
