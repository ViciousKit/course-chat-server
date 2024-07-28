package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
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
	storage *storage.Storage
}

const (
	errorMissingArguments = "missing arguments"
	errorInternal         = "internal error"
	errorMissingEntity    = "missing requested entity"
)

func main() {
	cfg := config.LoadConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Started app at port :%d", cfg.GRPC.Port)

	api := &srv{}
	connection := initStorage(cfg.PGUsername, cfg.PGPassword, cfg.PGDatabase, cfg.PGHost, cfg.PGPort)
	api.storage = storage.New(connection)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	generated.RegisterChatServerV1Server(grpcServer, api)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func initStorage(user string, password string, dbname string, host string, port int) *pgx.Conn {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Println("Cant connect pg" + err.Error())
		panic(err)
	}
	if err := conn.Ping(context.Background()); err != nil {
		fmt.Println("Cant ping pg" + err.Error())
		panic(err)
	}
	fmt.Println("Connected!")

	return conn
}

func (s *srv) SendMessage(ctx context.Context, req *generated.SendMessageRequest) (*emptypb.Empty, error) {
	if req.GetChatId() == 0 || req.GetFrom() == 0 || req.GetText() == "" || req.GetTimestamp() == nil {
		return &emptypb.Empty{}, fmt.Errorf(errorMissingArguments)
	}

	err := s.storage.SendMessage(ctx, req.GetFrom(), req.GetText(), req.GetChatId(), req.GetTimestamp())
	if err != nil {
		fmt.Println(err)

		return &emptypb.Empty{}, fmt.Errorf(errorInternal)
	}

	return &emptypb.Empty{}, nil
}

func (s *srv) CreateChat(ctx context.Context, req *generated.CreateChatRequest) (*generated.CreateChatResponse, error) {
	if len(req.UserIds) == 0 {
		return &generated.CreateChatResponse{}, fmt.Errorf(errorMissingArguments)
	}
	chatId, err := s.storage.CreateChat(ctx, req.GetUserIds())
	if err != nil {
		fmt.Println(err)

		return &generated.CreateChatResponse{}, fmt.Errorf(errorInternal)
	}

	return &generated.CreateChatResponse{
		Id: chatId,
	}, nil
}

func (s *srv) DeleteChat(ctx context.Context, req *generated.DeleteChatRequest) (*emptypb.Empty, error) {
	if req.GetId() == 0 {
		return &emptypb.Empty{}, fmt.Errorf(errorMissingArguments)
	}

	err := s.storage.DeleteChat(ctx, req.GetId())
	if err != nil {
		fmt.Println(err)

		return &emptypb.Empty{}, fmt.Errorf(errorInternal)
	}

	return &emptypb.Empty{}, nil
}

func (s *srv) AddUsers(ctx context.Context, req *generated.AddUsersRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *srv) RemoveUsers(ctx context.Context, req *generated.RemoveUsersRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *srv) GetMessages(ctx context.Context, req *generated.GetMessagesRequest) (*generated.GetMessagesResponse, error) {
	return &generated.GetMessagesResponse{}, nil
}
