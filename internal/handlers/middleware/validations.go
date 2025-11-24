package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
	"net/http"
	"regexp"
)

func ValidateID(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			logger.Warn("invalid id", zap.String("id", id))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			c.Abort()
			return
		}

		c.Set("id", id)
		c.Next()
	}
}

func ValidateNewSubscription(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.SubscriptionRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			logger.Warn("invalid request body", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body " + err.Error(),
			})
			c.Abort()
			return
		}

		valid := validator.New()
		if err := valid.Struct(req); err != nil {
			var validationErrors []string
			for _, err := range err.(validator.ValidationErrors) {
				field := err.Field()
				tag := err.Tag()

				switch tag {
				case "required":
					validationErrors = append(validationErrors, fmt.Sprintf("%s is required", field))
				case "gt":
					validationErrors = append(validationErrors, fmt.Sprintf("%s must be greater then 0", field))
				case "oneof":
					validationErrors = append(validationErrors, fmt.Sprintf("%s must be one of [1, 3, 6, 12]", field))
				default:
					validationErrors = append(validationErrors, fmt.Sprintf("%s failed validation: %s", field, tag))
				}
			}

			logger.Warn("validation failed", zap.Any("errors", validationErrors))
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "validation failed",
				"details": validationErrors,
			})
			c.Abort()
			return
		}

		logger.Info("Validated operation",
			zap.String("service_name", req.ServiceName),
			zap.Uint("price", req.Price),
			zap.String("user_id", req.UserID),
			zap.String("start_date", req.StartDate),
			zap.Uint("period", req.Period),
		)

		c.Set("payload_subscription", req)
		c.Next()
	}
}

func ValidateGetTotalCost(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params domain.TotalCostParams

		if err := c.ShouldBindQuery(&params); err != nil {
			logger.Warn("Invalid query params", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query params"})
			c.Abort()
			return
		}

		valid := validator.New(validator.WithRequiredStructEnabled())

		if err := valid.RegisterValidation("monthyear", validateMonthYear); err != nil {
			logger.Warn("Failed validate start or end date", zap.Error(err))
		}

		if err := valid.Struct(&params); err != nil {
			var validationErrors []string
			for _, err := range err.(validator.ValidationErrors) {
				field := err.Field()
				tag := err.Tag()

				switch tag {
				case "required":
					validationErrors = append(validationErrors, fmt.Sprintf("%s is required", field))
				case "len":
					validationErrors = append(validationErrors, fmt.Sprintf("%s must be 7 characters (MM-YYYY)", field))
				case "monthyear":
					validationErrors = append(validationErrors, fmt.Sprintf("%s must be format MM-YYYY (example: 01-2025)", field))
				case "gtefield":
					validationErrors = append(validationErrors, fmt.Sprintf("end_date cannot be earlier than start_date"))
				default:
					validationErrors = append(validationErrors, fmt.Sprintf("%s failed validation: %s", field, tag))
				}
			}

			logger.Warn("validation failed", zap.Any("errors", validationErrors))
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "validation failed",
				"details": validationErrors})
			c.Abort()
			return
		}

		c.Set("total_cost_params", params)
		c.Next()
	}
}

func validateMonthYear(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])-\d{4}$`)
	return re.MatchString(value)
}
