package subscription

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
)

func (r *repository) GetTotalCost(ctx context.Context, userID string, params domain.TotalCostWithPeriod) (uint, error) {
	r.logger.Info("Calculating total subscription cost for period")

	queryStr := sq.Select("COALESCE(SUM(price), 0)").From("subscriptions").
		Where(sq.Eq{"user_id": userID})

	if params.ServiceName != "" {
		queryStr = queryStr.Where(sq.Eq{"service_name": params.ServiceName})
	}

	queryStr = queryStr.Where("start_date <= ?", params.Period.End)
	queryStr = queryStr.Where(sq.Or{
		sq.Eq{"end_date": nil},
		sq.Expr("end_date >= ?", params.Period.Start),
	})

	query, args, err := queryStr.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build total cost query", zap.Error(err))
		return 0, err
	}

	var total uint
	err = r.db.QueryRow(ctx, query, args...).Scan(&total)
	if err != nil {
		r.logger.Error("Failed to calculate total cost", zap.Error(err))
		return 0, err
	}

	return total, nil
}
