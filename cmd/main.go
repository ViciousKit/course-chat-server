package main

import (
	"context"
	"fmt"
	"log"
	"net"

	generated "github.com/ViciousKit/course-chat-server/generated/chat_server_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type srv struct {
	generated.UnimplementedChatServerV1Server
}

const port = 8085

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Started app at port :%d", port)

	api := &srv{}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	generated.RegisterChatServerV1Server(grpcServer, api)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func (*srv) Create(ctx context.Context, req *generated.CreateRequest) (*generated.CreateResponse, error) {
	fmt.Println("__Create__")
	fmt.Println(req.GetUserNames())

	return &generated.CreateResponse{}, nil
}

func (*srv) Delete(ctx context.Context, req *generated.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Println("__Delete__")

	return &emptypb.Empty{}, nil
}

func (*srv) SendMessage(ctx context.Context, req *generated.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Println("__SendMessage__")

	return &emptypb.Empty{}, nil
}
