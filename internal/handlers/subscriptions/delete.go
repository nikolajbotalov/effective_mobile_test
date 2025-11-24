package subscriptions

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	"go.uber.org/zap"
	"net/http"
)

// DeleteSubscription godoc
// @Summary Удаление подписки
// @Description Удаляет подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription_id path string true "ID подписки (UUID)"
// @Success 200 {string} string
// @Failure 404 {object} map[string]string "Невалидный ID"
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [delete]
func DeleteSubscription(uc subscriptions.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Deleting subscription")

		id, exists := c.Get("id")
		if !exists {
			logger.Warn("Subscription id required")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription id required"})
			return
		}

		idParam, ok := id.(string)
		if !ok {
			logger.Warn("invalid subscription id type")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid subscription id type"})
			return
		}

		if err := uc.DeleteSubscription(c.Request.Context(), idParam); err != nil {
			logger.Error("Failed to delete subscription", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subscription"})
			return
		}

		logger.Info("Successfully deleted subscription", zap.String("id", idParam))
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
