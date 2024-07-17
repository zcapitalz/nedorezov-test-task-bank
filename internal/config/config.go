package config

import (
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel   string           `env:"LOG_LEVEL" env-default:"info"`
	HTTPServer HTTPServerConfig `env-prefix:"HTTP_SERVER_"`
}

type HTTPServerConfig struct {
	IpAddress string        `env:"IP_ADDRESS" env-default:"localhost"`
	Port      string        `env:"PORT" env-default:"1312"`
	Timeout   time.Duration `env:"TIMEOUT" env-default:"4s"`
}

var (
	once sync.Once
	cfg  Config
)

func MustNew() Config {
	once.Do(func() {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			log.Fatalf("could not read config: %s", err)
		}
	})

	return cfg
}
