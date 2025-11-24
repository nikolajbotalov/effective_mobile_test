package subscription

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (r *repository) UpdateSubscription(ctx context.Context, s *domain.Subscription) (*domain.Subscription, error) {
	r.logger.Info("Updating subscription", zap.String("id", s.ID))

	query, args, err := sq.Update("subscriptions").
		Set("service_name", s.ServiceName).
		Set("price", s.Price).
		Set("user_id", s.UserID).
		Set("start_date", s.StartDate).
		Set("end_date", s.EndDate).
		Set("updated_at", sq.Expr("NOW()")).
		Where(sq.Eq{"id": s.ID}).
		Suffix("RETURNING id, service_name, price, user_id, start_date, end_date, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("failed to build update query", zap.Error(err))
		return nil, err
	}

	var updatedSubscription domain.Subscription
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&updatedSubscription.ID,
		&updatedSubscription.ServiceName,
		&updatedSubscription.Price,
		&updatedSubscription.UserID,
		&updatedSubscription.StartDate,
		&updatedSubscription.EndDate,
		&updatedSubscription.CreatedAt,
		&updatedSubscription.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		r.logger.Warn("subscription not found", zap.String("id", s.ID))
		return nil, domain.ErrSubscriptionNotFound
	}

	if err != nil {
		r.logger.Error("failed to update subscription", zap.Error(err))
		return nil, err
	}

	r.logger.Info("Successfully updated subscription", zap.String("id", s.ID))
	return &updatedSubscription, nil
}
