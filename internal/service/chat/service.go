package chat

import (
	"context"
	"github.com/ViciousKit/course-chat-server/internal/repository"
	"github.com/ViciousKit/course-chat-server/internal/service"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Service struct {
	repository repository.ChatRepository
}

func NewService(r repository.ChatRepository) service.ChatService {
	return &Service{repository: r}
}

func (s *Service) CreateChat(ctx context.Context, userIds []int64) (int64, error) {
	return s.repository.CreateChat(ctx, userIds)
}
func (s *Service) DeleteChat(ctx context.Context, chatId int64) error {
	return s.repository.DeleteChat(ctx, chatId)
}
func (s *Service) SendMessage(ctx context.Context, from int64, text string, chatId int64, timestamp *timestamp.Timestamp) error {
	return s.repository.SendMessage(ctx, from, text, chatId, timestamp)
}
func (s *Service) AddUsers(ctx context.Context, userIds []int64, chatId int64) error {
	return s.repository.AddUsers(ctx, userIds, chatId)
}
func (s *Service) RemoveUsers(ctx context.Context, userIds []int64, chatId int64) error {
	return s.repository.RemoveUsers(ctx, userIds, chatId)
}
func (s *Service) GetMessages(ctx context.Context, chatId int64, lastId int64, limit int64, offset int64) error {
	return s.repository.GetMessages(ctx, chatId, lastId, limit, offset)
}
