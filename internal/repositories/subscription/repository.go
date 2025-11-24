package subscription

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

type Repository interface {
	GetSubscriptions(ctx context.Context, params domain.GetAllSubscriptionsParams) ([]domain.Subscription, error)
	GetByID(ctx context.Context, id string) (*domain.Subscription, error)
	CreateSubscription(ctx context.Context, s *domain.Subscription) error
	UpdateSubscription(ctx context.Context, s *domain.Subscription) (*domain.Subscription, error)
	DeleteSubscription(ctx context.Context, id string) error
	GetTotalCost(ctx context.Context, userID string, params domain.TotalCostWithPeriod) (uint, error)
}

type repository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewRepository(db *pgxpool.Pool, logger *zap.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
