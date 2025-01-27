package app

import (
	"context"
	"fmt"
	"github.com/ViciousKit/course-chat-server/internal/service"
	"log"

	"github.com/ViciousKit/course-chat-server/internal/api"
	chatApi "github.com/ViciousKit/course-chat-server/internal/api"
	"github.com/ViciousKit/course-chat-server/internal/closer"
	"github.com/ViciousKit/course-chat-server/internal/config"
	"github.com/ViciousKit/course-chat-server/internal/repository"
	chatRepository "github.com/ViciousKit/course-chat-server/internal/repository/chat"
	chatService "github.com/ViciousKit/course-chat-server/internal/service/chat"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceProvider struct {
	PGConfig   *config.PGConfig
	GRPCConfig *config.GRPCConfig

	pgPool *pgxpool.Pool

	chatRepository repository.ChatRepository

	chatService service.ChatService

	api *api.Srv
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) GetPGConfig() config.PGConfig {
	if sp.PGConfig == nil {
		cfg, err := config.LoadPGConfig()
		if err != nil {
			panic(err)
		}
		sp.PGConfig = &cfg
	}
	return *sp.PGConfig
}

func (sp *ServiceProvider) GetGRPCConfig() config.GRPCConfig {
	if sp.GRPCConfig == nil {
		cfg, err := config.LoadGRPCConfig()
		if err != nil {
			panic(err)
		}
		sp.GRPCConfig = &cfg
	}
	return *sp.GRPCConfig
}

func (sp *ServiceProvider) GetPGPool(ctx context.Context) *pgxpool.Pool {
	if sp.pgPool == nil {
		pool, err := pgxpool.New(ctx, sp.GetPGConfig().PGDSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		if err := pool.Ping(ctx); err != nil {
			fmt.Println("Cant ping pg" + err.Error())
			panic(err)
		}
		fmt.Println("PG Connected!")

		closer.Add(func() error {
			pool.Close()

			return nil
		})

		sp.pgPool = pool
	}

	return sp.pgPool
}

func (sp *ServiceProvider) GetChatRepository(ctx context.Context) repository.ChatRepository {
	if sp.chatRepository == nil {
		sp.chatRepository = chatRepository.NewRepository(sp.GetPGPool(ctx))
	}
	return sp.chatRepository
}

func (sp *ServiceProvider) GetChatService(ctx context.Context) service.ChatService {
	if sp.chatService == nil {
		sp.chatService = chatService.NewService(sp.GetChatRepository(ctx))
	}
	return sp.chatService
}

func (sp *ServiceProvider) GetChatApi(ctx context.Context) *api.Srv {
	if sp.api == nil {
		sp.api = chatApi.NewApi(sp.GetChatService(ctx))
	}
	return sp.api
}
