package app

import (
	"context"
	"fmt"
	generated "github.com/ViciousKit/course-chat-server/generated/chat_server_v1"
	"github.com/ViciousKit/course-chat-server/internal/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {
	grpcServer      *grpc.Server
	serviceProvider *ServiceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDependencies(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) initDependencies(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run() error {
	fmt.Println("Running app")

	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = NewServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)
	generated.RegisterChatServerV1Server(grpcServer, a.serviceProvider.GetChatApi(ctx))

	a.grpcServer = grpcServer

	return nil
}

func (a *App) runGRPCServer() error {
	address := fmt.Sprintf(":%s", a.serviceProvider.GetGRPCConfig().Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Started app at port :%s\n", address)
	if err := a.grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

	return nil
}
