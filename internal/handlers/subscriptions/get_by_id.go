package subscriptions

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"github.com/nikolajbotalov/effective_mobile_test/internal/handlers/helpers"
	"github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	"go.uber.org/zap"
	"net/http"
)

// GetByID godoc
// @Summary Получить подписку по ID
// @Description Возвращает подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription_id path string true "ID подписки (UUID)"
// @Success 200 {object} domain.Subscription
// @Failure 400 {object} map[string]string "Невалидный ID"
// @Failure 404 {string} string "Подписка не найдена"
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [get]
func GetByID(uc subscriptions.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Getting subscription")

		id, err := helpers.GetContextID(c, logger)
		if err != nil {
			logger.Error("ID is required", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ID is required"})
			return
		}

		subscription, err := uc.GetByID(c.Request.Context(), id)
		if err != nil {
			logger.Error("Failed to get subscription", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription"})
			return
		}

		logger.Info("Successfully got subscription", zap.String("id", id))
		c.JSON(http.StatusOK, domain.Subscription{
			ID:          subscription.ID,
			ServiceName: subscription.ServiceName,
			Price:       subscription.Price,
			UserID:      subscription.UserID,
			StartDate:   subscription.StartDate,
			EndDate:     subscription.EndDate,
			CreatedAt:   subscription.CreatedAt,
			UpdatedAt:   subscription.UpdatedAt,
		})
	}
}
