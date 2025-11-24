package subscriptions

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"github.com/nikolajbotalov/effective_mobile_test/internal/handlers/helpers"
	"github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	"go.uber.org/zap"
	"net/http"
)

// UpdateSubscription godoc
// @Summary Обновление подписки
// @Description Обновляет подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription_id path string true "ID подписки (UUID)"
// @Param subscription body domain.SubscriptionRequest true "Данные подписки"
// @Success 200 {object} domain.Subscription
// @Failure 400 {object} map[string]string "Невалидные данные"
// @Failure 404 {string} string "Подписка не найдена"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [put]
func UpdateSubscription(uc subscriptions.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Updating subscription")

		id, err := helpers.GetContextID(c, logger)
		if err != nil {
			logger.Error("ID is required", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ID is required"})
			return
		}

		req, exists := c.Get("payload_subscription")
		if !exists {
			logger.Error("did not get payload_subscription")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		subReq := req.(domain.SubscriptionRequest)

		subRes, err := uc.UpdateSubscription(c.Request.Context(), id, subReq)
		if err != nil {
			logger.Error("failed to update subscription", zap.Error(err))

			if errors.Is(err, domain.ErrSubscriptionNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		logger.Info("Successfully updated subscription", zap.String("id", id))
		c.JSON(http.StatusOK, domain.Subscription{
			ID:          subRes.ID,
			ServiceName: subRes.ServiceName,
			Price:       subRes.Price,
			UserID:      subRes.UserID,
			StartDate:   subRes.StartDate,
			EndDate:     subRes.EndDate,
			CreatedAt:   subRes.CreatedAt,
			UpdatedAt:   subRes.UpdatedAt,
		})
	}
}
