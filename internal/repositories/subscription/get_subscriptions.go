package subscription

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (r *repository) GetSubscriptions(ctx context.Context, params domain.GetAllSubscriptionsParams) ([]domain.Subscription, error) {
	r.logger.Info("Fetching subscriptions")

	query, args, err := sq.Select("*").From("subscriptions").Limit(uint64(params.Limit)).
		Offset(uint64(params.Offset)).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build query", zap.Error(err))
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		r.logger.Error("Failed to execute query", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var subscriptions []domain.Subscription
	for rows.Next() {
		var subscription domain.Subscription
		if err = rows.Scan(
			&subscription.ID,
			&subscription.ServiceName,
			&subscription.Price,
			&subscription.UserID,
			&subscription.StartDate,
			&subscription.EndDate,
			&subscription.CreatedAt,
			&subscription.UpdatedAt,
		); err != nil {
			r.logger.Error("Failed to scan row", zap.Error(err))
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error while iterating over rows", zap.Error(err))
		return nil, err
	}

	if len(subscriptions) == 0 {
		r.logger.Info("No subscriptions found")
		return subscriptions, nil
	}

	r.logger.Info("Successfully fetched subscriptions", zap.Int("count", len(subscriptions)))

	return subscriptions, nil
}
