package subscriptions

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	"go.uber.org/zap"
	"net/http"
)

// CreateSubscription godoc
// @Summary Добавить подписку
// @Description Добавляет новую подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body domain.SubscriptionRequest true "Данные подписки"
// @Success 201 {string} string
// @Failure 400 {object} map[string]string "Невалидные данные"
// @Failure 409 {string} map[string]string "Подписка уже существует"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /subscriptions [post]
func CreateSubscription(uc subscriptions.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Creating new subscription")
		req, exists := c.Get("payload_subscription")
		if !exists {
			logger.Error("did not get payload_subscription")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		subReq := req.(domain.SubscriptionRequest)

		err := uc.CreateSubscription(c.Request.Context(), subReq)

		if err != nil {
			logger.Error("failed to create new subscription", zap.Error(err))

			if errors.Is(err, domain.ErrSubscriptionAlreadyExists) {
				c.JSON(http.StatusConflict, gin.H{
					"error": "user already has subscription for this service",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		logger.Info("Successfully created new subscription")
		c.JSON(http.StatusCreated, gin.H{"status": "success"})
	}
}
