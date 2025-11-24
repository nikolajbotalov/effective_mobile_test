package subscriptions

import (
	"context"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (uc *useCase) GetTotalCost(ctx context.Context, userID string, params domain.TotalCostWithPeriod) (uint, error) {
	uc.logger.Info("Processing get total cost", zap.String("user_id", userID))

	total, err := uc.repo.GetTotalCost(ctx, userID, params)
	if err != nil {
		uc.logger.Error("Failed to calculate total cost", zap.Error(err))
		return 0, err
	}

	uc.logger.Info("Successfully processing get total cost", zap.Uint("total_cost", total))
	return total, nil
}
