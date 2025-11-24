package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/domain"
	"go.uber.org/zap"
	"net/http"
)

func PaginationMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Parsing pagination params")
		var params domain.GetAllSubscriptionsParams

		if err := c.ShouldBindQuery(&params); err != nil {
			logger.Warn("Invalid pagination params", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit or offset"})
			c.Abort()
			return
		}

		fmt.Printf("params: %+v\n", params)

		if params.Limit <= 0 {
			params.Limit = 10
		}
		if params.Offset < 0 {
			params.Offset = 0
		}

		c.Set("pagination", params)
		c.Next()
	}
}
