package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const defaultConfigPath = "local.env"

type Config struct {
	GRPC       GRPCConfig `env:"GRPC"`
	PGUsername string     `env:"PG_USER"`
	PGPassword string     `env:"PG_PASSWORD"`
	PGDatabase string     `env:"PG_DATABASE_NAME"`
	PGHost     string     `env:"PG_HOST"`
	PGPort     int        `env:"PG_PORT"`
}

type GRPCConfig struct {
	Port int    `env:"GRPC_PORT" envDefault:"50051"`
	Host string `env:"GRPC_HOST" envDefault:"localhost"`
}

func LoadConfig() Config {
	var cfgPath string

	flag.StringVar(&cfgPath, "config", "", "path to cfg file")
	flag.Parse()

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Println(e.Name())
	}

	if cfgPath == "" {
		cfgPath = defaultConfigPath
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic("config file does not exist: " + cfgPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return cfg
}
