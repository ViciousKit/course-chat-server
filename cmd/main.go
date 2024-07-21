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
	return &emptypb.Empty{}, nil
}

func (s *srv) CreateChat(ctx context.Context, req *generated.CreateChatRequest) (*generated.CreateChatResponse, error) {
	return &generated.CreateChatResponse{}, nil
}

func (s *srv) DeleteChat(ctx context.Context, req *generated.DeleteChatRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *srv) AddUsers(ctx context.Context, req *generated.AddUsersRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *srv) RemoveUsers(ctx context.Context, req *generated.RemoveUsersRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
