package main

import (
	"context"
	"fmt"
	"log"
	"net"

	generated "github.com/ViciousKit/course-chat-server/generated/chat_server_v1"
	"github.com/ViciousKit/course-chat-server/internal/config"
	"github.com/ViciousKit/course-chat-server/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type srv struct {
	generated.UnimplementedChatServerV1Server
	Storage *storage.Storage
}

const (
	errorMissingArguments    = "missing arguments"
	errorInternal            = "internal error"
	errorPasswordDoesntMatch = "password doesn't match"
	errorMissingEntity       = "missing requested entity"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Printf("%+v\n", cfg)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Started app at port :%d", cfg.GRPC.Port)

	api := &srv{}
	api.Storage = storage.New(cfg.PGUsername, cfg.PGPassword, cfg.PGDatabase, cfg.PGHost, cfg.PGPort)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	generated.RegisterChatServerV1Server(grpcServer, api)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func (s *srv) SendMessage(ctx context.Context, req *generated.SendMessageRequest) (*emptypb.Empty, error) {
	method := "SendMessage"

	if req.GetChatId() == 0 || req.GetFrom() == 0 || req.GetText() == "" || req.GetTimestamp() == nil {
		return &emptypb.Empty{}, fmt.Errorf("%s: %s", method, errorMissingArguments)
	}

	err := s.Storage.SendMessage(ctx, req.GetFrom(), req.GetText(), req.GetChatId(), req.GetTimestamp())
	if err != nil {
		fmt.Println(err)

		return &emptypb.Empty{}, fmt.Errorf("%s: %s", method, errorInternal)
	}

	return &emptypb.Empty{}, nil
}

func (s *srv) CreateChat(ctx context.Context, req *generated.CreateChatRequest) (*generated.CreateChatResponse, error) {
	method := "CreateChat"
	if len(req.UserIds) == 0 {
		return &generated.CreateChatResponse{}, fmt.Errorf("%s: %s", method, errorMissingArguments)
	}
	chatId, err := s.Storage.CreateChat(ctx, req.GetUserIds())
	if err != nil {
		fmt.Println(err)

		return &generated.CreateChatResponse{}, fmt.Errorf("%s: %s", method, errorInternal)
	}

	return &generated.CreateChatResponse{
		Id: chatId,
	}, nil
}

func (s *srv) DeleteChat(ctx context.Context, req *generated.DeleteChatRequest) (*emptypb.Empty, error) {
	method := "DeleteChat"

	if req.GetId() == 0 {
		return &emptypb.Empty{}, fmt.Errorf("%s: %s", method, errorMissingArguments)
	}

	err := s.Storage.DeleteChat(ctx, req.GetId())
	if err != nil {
		fmt.Println(err)

		return &emptypb.Empty{}, fmt.Errorf("%s: %s", method, errorInternal)
	}

	return &emptypb.Empty{}, nil
}

func (s *srv) AddUsers(ctx context.Context, req *generated.AddUsersRequest) (*emptypb.Empty, error) {
	// method := "AddUsers"
	return &emptypb.Empty{}, nil
}

func (s *srv) RemoveUsers(ctx context.Context, req *generated.RemoveUsersRequest) (*emptypb.Empty, error) {
	// method := "RemoveUsers"
	return &emptypb.Empty{}, nil
}

func (s *srv) GetMessages(ctx context.Context, req *generated.GetMessagesRequest) (*generated.GetMessagesResponse, error) {
	// method := "GetMessages"
	return &generated.GetMessagesResponse{}, nil
}
