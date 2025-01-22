package main

import (
	"context"
	"fmt"
	chatApi "github.com/ViciousKit/course-chat-server/internal/client/api"
	chatRepository "github.com/ViciousKit/course-chat-server/internal/repository/chat"
	chatService "github.com/ViciousKit/course-chat-server/internal/service/chat"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net"

	generated "github.com/ViciousKit/course-chat-server/generated/chat_server_v1"
	"github.com/ViciousKit/course-chat-server/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.LoadConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Started app at port :%d\n", cfg.GRPC.Port)

	pool := initStorage(cfg.PGUsername, cfg.PGPassword, cfg.PGDatabase, cfg.PGHost, cfg.PGPort)
	repository := chatRepository.NewRepository(pool)
	chatServ := chatService.NewService(repository)
	api := chatApi.NewApi(chatServ)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	generated.RegisterChatServerV1Server(grpcServer, api)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func initStorage(user string, password string, dbname string, host string, port int) *pgxpool.Pool {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		fmt.Println("Cant ping pg" + err.Error())
		panic(err)
	}
	fmt.Println("Connected!")

	return pool
}
