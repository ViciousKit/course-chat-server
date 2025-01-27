package service

import (
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatService interface {
	CreateChat(ctx context.Context, userIds []int64) (int64, error)
	DeleteChat(ctx context.Context, chatId int64) error
	SendMessage(ctx context.Context, from int64, text string, chatId int64, timestamp *timestamppb.Timestamp) error
	AddUsers(ctx context.Context, userIds []int64, chatId int64) error
	RemoveUsers(ctx context.Context, userIds []int64, chatId int64) error
	GetMessages(ctx context.Context, chatId int64, lastId int64, limit int64, offset int64) error
}
