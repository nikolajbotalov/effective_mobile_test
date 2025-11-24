package subscriptions

import (
	"context"
	"errors"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (uc *useCase) DeleteSubscription(ctx context.Context, id string) error {
	uc.logger.Info("Processing delete subscription", zap.String("id", id))

	if err := uc.repo.DeleteSubscription(ctx, id); err != nil {
		if errors.Is(err, domain.ErrSubscriptionNotFound) {
			uc.logger.Warn("Subscription not found", zap.String("id", id))
			return domain.ErrSubscriptionNotFound
		}
		return domain.ErrDeleted
	}

	uc.logger.Info("Successfully deleted subscription", zap.String("id", id))
	return nil
}
