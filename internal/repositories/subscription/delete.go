package subscription

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (r *repository) DeleteSubscription(ctx context.Context, id string) error {
	r.logger.Info("Delete subscription")

	query, args, err := sq.Delete("subscriptions").Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build query", zap.Error(err))
		return errors.New("failed to build delete query")
	}

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		r.logger.Error("Failed to delete subscription", zap.Error(err))
		return domain.ErrDeleted
	}

	if res.RowsAffected() == 0 {
		r.logger.Error("Subscription not found", zap.String("id", id))
		return errors.New("subscription not found")
	}

	r.logger.Info("Successfully deleted subscription", zap.String("id", id))
	return nil
}
