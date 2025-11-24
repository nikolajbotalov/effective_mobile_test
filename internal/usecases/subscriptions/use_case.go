package subscriptions

import (
	"context"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"github.com/nikolajbotalov/effective_mobile_test/internal/repositories/subscription"
	"go.uber.org/zap"
)

type UseCase interface {
	GetSubscriptions(ctx context.Context, params domain.GetAllSubscriptionsParams) ([]domain.Subscription, error)
	GetByID(ctx context.Context, id string) (*domain.Subscription, error)
	CreateSubscription(ctx context.Context, req domain.SubscriptionRequest) error
	UpdateSubscription(ctx context.Context, id string, req domain.SubscriptionRequest) (*domain.Subscription, error)
	DeleteSubscription(ctx context.Context, id string) error
	GetTotalCost(ctx context.Context, userID string, params domain.TotalCostWithPeriod) (uint, error)
}

type useCase struct {
	repo   subscription.Repository
	logger *zap.Logger
}

func NewUseCase(repo subscription.Repository, logger *zap.Logger) UseCase {
	return &useCase{repo: repo, logger: logger}
}
