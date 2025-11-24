package subscriptions

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/handlers/middleware"
	"github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	"go.uber.org/zap"
)

func SetupRoutes(g *gin.Engine, uc subscriptions.UseCase, logger *zap.Logger) {
	routers := g.Group("api/v1/subscriptions")
	{
		routers.GET("/", middleware.PaginationMiddleware(logger), GetSubscriptions(uc, logger))
		routers.GET("/:id", middleware.ValidateID(logger), GetByID(uc, logger))
		routers.POST("/", middleware.ValidateNewSubscription(logger), CreateSubscription(uc, logger))
		routers.PATCH("/:id", middleware.ValidateID(logger), middleware.ValidateNewSubscription(logger),
			UpdateSubscription(uc, logger))
		routers.DELETE("/:id", middleware.ValidateID(logger), DeleteSubscription(uc, logger))
		routers.GET("/total_cost/:id", middleware.ValidateID(logger), middleware.ValidateGetTotalCost(logger),
			GetTotalCost(uc, logger))
	}
}
