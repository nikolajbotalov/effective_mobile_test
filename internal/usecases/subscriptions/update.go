package subscriptions

import (
	"context"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (uc *useCase) UpdateSubscription(ctx context.Context, id string, req domain.SubscriptionRequest) (*domain.Subscription, error) {
	uc.logger.Info("Processing update subscription")

	subscription, err := uc.TransformReqToSubscription(id, req)
	if err != nil {
		uc.logger.Error("failed to transform subscription", zap.Error(err))
		return nil, domain.ErrTransform
	}

	updatedSubscription, err := uc.repo.UpdateSubscription(ctx, subscription)
	if err != nil {
		uc.logger.Error("failed to update subscription", zap.Error(err))
		return nil, err
	}

	uc.logger.Info("Successfully updated subscription", zap.String("id", updatedSubscription.ID))
	return updatedSubscription, nil
}
