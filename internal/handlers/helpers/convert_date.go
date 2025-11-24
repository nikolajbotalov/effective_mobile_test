package helpers

import (
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
	"time"
)

func ConvertDate(startDate, endDate string, logger *zap.Logger) (domain.Period, error) {
	start, err := parseMonthYear(startDate)
	if err != nil {
		logger.Error("Failed to parse start date", zap.Error(err))
		return domain.Period{}, domain.ErrParseDate
	}

	end, err := parseMonthYear(endDate)
	if err != nil {
		logger.Error("Failed to parse end date", zap.Error(err))
		return domain.Period{}, domain.ErrParseDate
	}

	logger.Info("Successfully parsed start and end date",
		zap.Time("start_date", start),
		zap.Time("end_date", end))
	return domain.Period{
		Start: time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(end.Year(), end.Month(), getLastDay(end), 23, 59, 59, 999999999, time.UTC),
	}, nil
}

func parseMonthYear(s string) (time.Time, error) {
	return time.Parse("01-2006", s)
}

func getLastDay(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
