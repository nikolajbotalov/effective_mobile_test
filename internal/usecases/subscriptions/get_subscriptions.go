package subscriptions

import (
	"context"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (uc *useCase) GetSubscriptions(ctx context.Context, params domain.GetAllSubscriptionsParams) ([]domain.Subscription, error) {
	uc.logger.Info("Processing get subscriptions", zap.Int("limit", params.Limit), zap.Int("offset", params.Offset))

	if err := ValidatePagination(params.Limit, params.Offset); err != nil {
		uc.logger.Warn("Invalid pagination params", zap.Int("limit", params.Limit), zap.Int("offset", params.Offset))
		return nil, err
	}

	subscriptions, err := uc.repo.GetSubscriptions(ctx, params)
	if err != nil {
		uc.logger.Error("Failed to get subscriptions", zap.Error(err))
		return nil, err
	}

	uc.logger.Info("Successfully processed get subscriptions", zap.Int("subscriptions count", len(subscriptions)))
	return subscriptions, nil
}
