package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"os"
	"sync"
)

type Config struct {
	Listen        Listen
	DB            DB
	RetryAttempts uint `env:"RETRY_ATTEMPTS" envDefault:"10"`
}

type Listen struct {
	BindIP string `env:"BIND_IP" envDefault:"0.0.0.0"`
	Port   string `env:"PORT" envDefault:"8080"`
}

type DB struct {
	User     string `env:"PG_USER" envDefault:"postgres"`
	Password string `env:"PG_PASSWORD" envDefault:"admin"`
	Host     string `env:"PG_HOST" envDefault:"localhost"`
	Port     string `env:"PG_PORT" envDefault:"5432"`
	Name     string `env:"PG_NAME" envDefault:"subscriptions"`
}

var instance *Config
var once sync.Once

func LoadConfig(logger *zap.Logger) *Config {
	once.Do(func() {
		instance = &Config{}

		if _, err := os.Stat(".env"); err == nil {
			if err := cleanenv.ReadConfig(".env", instance); err != nil {
				logger.Error("failed loading .env file", zap.Error(err))
			}
		}

		if err := cleanenv.ReadEnv(instance); err != nil {
			logger.Error("failed loading config", zap.Error(err))
		}
	})

	logger.Info("config loaded")

	return instance
}
