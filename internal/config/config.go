package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type PGConfig struct {
	GRPC       GRPCConfig `env:"GRPC"`
	PGUsername string     `env:"PG_USER"`
	PGPassword string     `env:"PG_PASSWORD"`
	PGDatabase string     `env:"PG_DATABASE_NAME"`
	PGHost     string     `env:"PG_HOST"`
	PGPort     string     `env:"PG_PORT"`
}

type GRPCConfig struct {
	Port string `env:"GRPC_PORT" envDefault:"50051"`
	Host string `env:"GRPC_HOST" envDefault:"localhost"`
}

type Path struct {
	path string
}

var configPath = New()

func New() *Path {
	p := &Path{}
	_ = p.Load()

	return p
}

func LoadPGConfig() (PGConfig, error) {
	if _, err := os.Stat(configPath.path); os.IsNotExist(err) {
		return PGConfig{}, errors.New(fmt.Sprintf("config file does not exist: " + configPath.path))
	}

	var cfg PGConfig
	if err := cleanenv.ReadConfig(configPath.path, &cfg); err != nil {
		return PGConfig{}, errors.New(fmt.Sprintf("failed to read pg config: " + err.Error()))
	}

	return cfg, nil
}

func LoadGRPCConfig() (GRPCConfig, error) {
	if _, err := os.Stat(configPath.path); os.IsNotExist(err) {
		return GRPCConfig{}, errors.New(fmt.Sprintf("config file does not exist: %s\n", configPath.path))
	}

	var cfg GRPCConfig
	if err := cleanenv.ReadConfig(configPath.path, &cfg); err != nil {
		return GRPCConfig{}, errors.New(fmt.Sprintf("failed to read grpc config: " + err.Error()))
	}

	return cfg, nil
}

func (cfg PGConfig) PGDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.PGHost, cfg.PGPort, cfg.PGUsername, cfg.PGPassword, cfg.PGDatabase)
}
