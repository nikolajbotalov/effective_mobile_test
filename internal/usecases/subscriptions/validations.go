package subscriptions

import "github.com/nikolajbotalov/effective_mobile_test/internal/domain"

func ValidatePagination(limit, offset int) error {
	if limit <= 0 || offset < 0 {
		return domain.ErrInvalidPagination
	}
	if limit > 100 {
		return domain.ErrLimitExceeded
	}
	return nil
}
