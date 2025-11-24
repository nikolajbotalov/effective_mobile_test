package subscriptions

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	"go.uber.org/zap"
	"net/http"
)

// GetSubscriptions godoc
// @Summary Получить список подписок
// @Description Возвращает список подписок
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param limit query int false "Количество записей на странице (по умолчанию = 10)"
// @Param offset query int false "Смещение (по умолчанию = 0)"
// @Success 200 {array} domain.GetAllSubscriptionsResponse "Список подписок с мета-данными"
// @Failure 400 {object} map[string]string "Некорректные параметры пагинации"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /subscriptions [get]
func GetSubscriptions(uc subscriptions.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Getting subscriptions")

		params, exists := c.Get("pagination")
		if !exists {
			logger.Warn("pagination params missing")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "pagination params missing"})
			return
		}

		paginationParams, ok := params.(domain.GetAllSubscriptionsParams)
		if !ok {
			logger.Warn("invalid pagination params type")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid pagination params type"})
			return
		}

		subs, err := uc.GetSubscriptions(c.Request.Context(), paginationParams)
		if err != nil {
			logger.Error("Failed to get subscriptions", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscriptions"})
			return
		}

		logger.Info("Successfully got subscriptions", zap.Int("subscriptions count", len(subs)))
		c.JSON(http.StatusOK, domain.GetAllSubscriptionsResponse{
			Subscriptions: subs,
			Meta: struct {
				Limit  int `json:"limit"`
				Offset int `json:"offset"`
				Total  int `json:"total"`
			}{Limit: paginationParams.Limit, Offset: paginationParams.Offset, Total: len(subs)},
		})
	}
}
