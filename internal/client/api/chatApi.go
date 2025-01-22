package api

import (
	"context"
	"fmt"
	generated "github.com/ViciousKit/course-chat-server/generated/chat_server_v1"
	"github.com/ViciousKit/course-chat-server/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Srv struct {
	generated.UnimplementedChatServerV1Server
	chatService service.ChatService
}

func NewApi(service service.ChatService) *Srv {
	return &Srv{
		chatService: service,
	}
}

const (
	errorMissingArguments = "missing arguments"
	errorInternal         = "internal error"
	errorMissingEntity    = "missing requested entity"
)

func (s *Srv) SendMessage(ctx context.Context, req *generated.SendMessageRequest) (*emptypb.Empty, error) {
	if req.GetChatId() == 0 || req.GetAuthorId() == 0 || req.GetText() == "" || req.GetTimestamp() == nil {
		return &emptypb.Empty{}, fmt.Errorf(errorMissingArguments)
	}

	err := s.chatService.SendMessage(ctx, req.GetAuthorId(), req.GetText(), req.GetChatId(), req.GetTimestamp())
	if err != nil {
		fmt.Println(err)

		return &emptypb.Empty{}, fmt.Errorf(errorInternal)
	}

	return &emptypb.Empty{}, nil
}

func (s *Srv) CreateChat(ctx context.Context, req *generated.CreateChatRequest) (*generated.CreateChatResponse, error) {
	if len(req.UserIds) == 0 {
		return &generated.CreateChatResponse{}, fmt.Errorf(errorMissingArguments)
	}
	chatId, err := s.chatService.CreateChat(ctx, req.GetUserIds())
	if err != nil {
		fmt.Println(err)

		return &generated.CreateChatResponse{}, fmt.Errorf(errorInternal)
	}

	return &generated.CreateChatResponse{
		Id: chatId,
	}, nil
}

func (s *Srv) DeleteChat(ctx context.Context, req *generated.DeleteChatRequest) (*emptypb.Empty, error) {
	if req.GetId() == 0 {
		return &emptypb.Empty{}, fmt.Errorf(errorMissingArguments)
	}

	err := s.chatService.DeleteChat(ctx, req.GetId())
	if err != nil {
		fmt.Println(err)

		return &emptypb.Empty{}, fmt.Errorf(errorInternal)
	}

	return &emptypb.Empty{}, nil
}

func (s *Srv) AddUsers(ctx context.Context, req *generated.AddUsersRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Srv) RemoveUsers(ctx context.Context, req *generated.RemoveUsersRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Srv) GetMessages(ctx context.Context, req *generated.GetMessagesRequest) (*generated.GetMessagesResponse, error) {
	return &generated.GetMessagesResponse{}, nil
}
