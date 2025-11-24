package domain

import (
	"errors"
	"time"
)

type Subscription struct {
	ID          string     `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       uint       `json:"price"`
	UserID      string     `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type GetAllSubscriptionsParams struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type GetAllSubscriptionsResponse struct {
	Subscriptions []Subscription `json:"subscriptions"`
	Meta          struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	} `json:"meta"`
}

type SubscriptionRequest struct {
	ServiceName string `json:"service_name" validate:"required"`
	Price       uint   `json:"price" validate:"required,gt=0"`
	UserID      string `json:"user_id" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	Period      uint   `json:"period" validate:"omitempty,oneof=1 3 6 12"`
}

type SubscriptionRequestWithPeriod struct {
	ServiceName string `json:"service_name" validate:"required"`
	Price       uint   `json:"price" validate:"required,gt=0"`
	UserID      string `json:"user_id" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	Period      uint   `json:"period" validate:"omitempty,oneof=1 3 6 12"`
}

type TotalCostParams struct {
	StartDate   string `form:"start_date" binding:"required" validate:"required,len=7,monthyear"`
	EndDate     string `form:"end_date" binding:"required" validate:"required,len=7,monthyear,gtefield=StartDate"`
	ServiceName string `form:"service_name" validate:"omitempty,max=100"`
}

type Period struct {
	Start time.Time
	End   time.Time
}

type TotalCostWithPeriod struct {
	ServiceName string
	Period      Period
}

var (
	ErrInvalidPagination         = errors.New("invalid pagination params")
	ErrLimitExceeded             = errors.New("limit exceeded maximum allowed value")
	ErrSubscriptionNotFound      = errors.New("subscription not found")
	ErrDeleted                   = errors.New("failed to delete subscription")
	ErrTransform                 = errors.New("failed to transform request to subscription")
	ErrSubscriptionAlreadyExists = errors.New("subscription already exists")
	ErrParseDate                 = errors.New("failed to parse date")
)
