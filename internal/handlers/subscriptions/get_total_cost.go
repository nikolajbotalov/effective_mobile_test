package subscriptions

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"github.com/nikolajbotalov/effective_mobile_test/internal/handlers/helpers"
	"github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	"go.uber.org/zap"
	"net/http"
)

// GetTotalCost godoc
// @Summary Получить сумму всех подписок пользователя
// @Description Возвращает сумму всех подписок пользователя за период и по названию сервис
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param user_id path string true "ID пользователя (UUID)"
// @Param start_date query string true "Начала периода (формат MM-YYYY, например, 01-2025)"
// @Param end_date query string true "Конец периода (формат MM-YYYY)"
// @Param service_name query string false "Фильт по названию сервиса"
// @Success 200 {object} map[string]uint "total_cost"
// @Failure 400 {object} map[string]string "Невалидные данные"
// @Failure 404 {string} string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /subscriptions/total_cost/{user_id} [get]
func GetTotalCost(uc subscriptions.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Getting total subscription cost")

		userID, err := helpers.GetContextID(c, logger)
		if err != nil {
			logger.Error("ID is required", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ID is required"})
			return
		}

		params, paramsExists := c.Get("total_cost_params")
		if !paramsExists {
			logger.Warn("Total cost params missing")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Total cost params missing"})
			return
		}

		totalCostParams, ok := params.(domain.TotalCostParams)
		if !ok {
			logger.Warn("Invalid total cost params")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid total cost params"})
			return
		}

		period, err := helpers.ConvertDate(totalCostParams.StartDate, totalCostParams.EndDate, logger)
		if err != nil {
			logger.Warn("Failed to dates", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert dates"})
			return
		}

		totalCostWithPeriod := domain.TotalCostWithPeriod{
			ServiceName: totalCostParams.ServiceName,
			Period:      period,
		}

		total, err := uc.GetTotalCost(c.Request.Context(), userID, totalCostWithPeriod)
		if err != nil {
			logger.Warn("Failed to get total cost", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total cost"})
			return
		}

		logger.Info("Successfully got total cost", zap.Uint("total_cost", total))
		c.JSON(http.StatusOK, gin.H{
			"total_cost": total,
		})
	}
}
