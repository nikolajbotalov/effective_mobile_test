package subscription

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
	"time"
)

func (r *repository) CreateSubscription(ctx context.Context, s *domain.Subscription) error {
	r.logger.Info("Creating new subscription")

	querySub, argsSub, err := sq.Select("1").From("subscriptions").
		Where(sq.Eq{"user_id": s.UserID, "service_name": s.ServiceName}).Limit(1).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("failed to build getting subscription query", zap.Error(err))
		return err
	}

	var rowFound int

	err = r.db.QueryRow(ctx, querySub, argsSub...).Scan(&rowFound)
	if err == nil {
		r.logger.Error("subscription already exists", zap.Error(err))
		return domain.ErrSubscriptionAlreadyExists
	}

	if !errors.Is(err, sql.ErrNoRows) {
		r.logger.Error("failed to check if subscription exists", zap.Error(err))
		return err
	}

	now := time.Now().Format(time.RFC3339)

	query, args, err := sq.Insert("subscriptions").
		Columns("id", "service_name", "price", "user_id", "start_date", "end_date", "created_at", "updated_at").
		Values(s.ID, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, now, now).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build query", zap.Error(err))
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		r.logger.Error("Failed to insert subscription", zap.Error(err))
		return err
	}

	r.logger.Info("Successfully created new subscription")
	return nil
}
