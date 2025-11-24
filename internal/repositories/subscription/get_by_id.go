package subscription

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (r *repository) GetByID(ctx context.Context, id string) (*domain.Subscription, error) {
	r.logger.Info("Fetching subscription by id", zap.String("id", id))

	query, args, err := sq.Select("*").From("subscriptions").Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build query", zap.Error(err))
		return nil, err
	}

	var subscription domain.Subscription

	if err = r.db.QueryRow(ctx, query, args...).Scan(
		&subscription.ID,
		&subscription.ServiceName,
		&subscription.Price,
		&subscription.UserID,
		&subscription.StartDate,
		&subscription.EndDate,
		&subscription.CreatedAt,
		&subscription.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("failed to find subscription by id", zap.String("id", id))
			return nil, domain.ErrSubscriptionNotFound
		}

		r.logger.Error("Failed to scan row", zap.Error(err))
		return nil, err
	}

	r.logger.Info("Successfully fetched subscription by id", zap.String("id", id))
	return &subscription, nil
}
