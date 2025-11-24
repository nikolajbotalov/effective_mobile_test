package app

import (
	"github.com/nikolajbotalov/effective_mobile_test/internal/adapters/db"
	"github.com/nikolajbotalov/effective_mobile_test/internal/config"
	"github.com/nikolajbotalov/effective_mobile_test/internal/logger"
	subsRepo "github.com/nikolajbotalov/effective_mobile_test/internal/repositories/subscription"
	subsUseCase "github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	"go.uber.org/zap"
)

type App struct {
	Logger *zap.Logger
	Server *Server
}

func NewApp() (*App, error) {
	zapLogger, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	cfg := config.LoadConfig(zapLogger)

	if err = db.RunMigrations(cfg, zapLogger); err != nil {
		zapLogger.Error("Failed to run migrations", zap.Error(err))
		return nil, err
	}

	dbInstance, err := db.NewDB(cfg, zapLogger)
	if err != nil {
		zapLogger.Error("Failed to initialize DB", zap.Error(err))
		return nil, err
	}

	subRepo := subsRepo.NewRepository(dbInstance.Pool(), zapLogger)
	subUseCase := subsUseCase.NewUseCase(subRepo, zapLogger)

	server := NewServer(cfg, subUseCase, zapLogger)

	return &App{
		Logger: zapLogger,
		Server: server,
	}, nil
}
