package subscriptions

import (
	"errors"
	"github.com/google/uuid"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"slices"
	"time"
)

func (uc *useCase) TransformReqToSubscription(id string, req domain.SubscriptionRequest) (*domain.Subscription, error) {
	uc.logger.Info("Transformed request to subscription")

	var subID string
	if id != "" {
		subID = id
	} else {
		subID = uuid.New().String()
	}

	parsedStart, err := time.Parse("01-2006", req.StartDate)
	if err != nil {
		return nil, errors.New("invalid start_date format, expected MM-YYYY")
	}

	startDate := time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 1, 0, 0, 0, time.UTC)

	periodMonths := req.Period
	if periodMonths == 0 {
		periodMonths = 1
	}

	if !slices.Contains([]uint{1, 3, 6, 12}, periodMonths) {
		return nil, errors.New("period must be 1, 3, 6 or 12 months")
	}

	var endDate *time.Time

	if periodMonths > 0 {
		end := parsedStart.AddDate(0, int(periodMonths), 0)
		lastDate := time.Date(end.Year(), end.Month(), 0, 23, 59, 59, 999999999, time.UTC)
		endDate = &lastDate
	}

	return &domain.Subscription{
		ID:          subID,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}
