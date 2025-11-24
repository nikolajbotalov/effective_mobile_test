package subscriptions

import (
	"context"
	"errors"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (uc *useCase) GetByID(ctx context.Context, id string) (*domain.Subscription, error) {
	uc.logger.Info("Processing request get by ID")

	subscription, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrSubscriptionNotFound) {
			uc.logger.Warn("subscription not found", zap.String("id", id))
			return nil, domain.ErrSubscriptionNotFound
		}
		uc.logger.Error("failed to get subscription", zap.String("id", id))
		return nil, err
	}

	uc.logger.Info("Successfully got subscription", zap.String("id", id))
	return subscription, nil
}
