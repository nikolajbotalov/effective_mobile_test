package subscriptions

import (
	"context"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (uc *useCase) CreateSubscription(ctx context.Context, req domain.SubscriptionRequest) error {
	uc.logger.Info("Processing create new subscriptions")

	subscription, err := uc.TransformReqToSubscription("", req)
	if err != nil {
		uc.logger.Error("failed to transform subscription", zap.Error(err))
		return domain.ErrTransform
	}

	if err = uc.repo.CreateSubscription(ctx, subscription); err != nil {
		uc.logger.Error("Failed to create new subscription", zap.Error(err))
		return err
	}

	uc.logger.Info("Successfully created new subscription")
	return nil
}
