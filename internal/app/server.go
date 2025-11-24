package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nikolajbotalov/effective_mobile_test/internal/config"
	subsHandler "github.com/nikolajbotalov/effective_mobile_test/internal/handlers/subscriptions"
	subsUseCase "github.com/nikolajbotalov/effective_mobile_test/internal/usecases/subscriptions"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
}

func NewServer(cfg *config.Config, useCase subsUseCase.UseCase, logger *zap.Logger) *Server {
	router := gin.New()

	router.Use(ginZapLogger(logger))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// TODO: написать корсы

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	subsHandler.SetupRoutes(router, useCase, logger)

	address := fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	return &Server{
		httpServer: &http.Server{
			Addr:    address,
			Handler: router,
		},
		logger: logger,
	}
}

func (s *Server) Run() {
	s.logger.Info("Starting server", zap.String("address", s.httpServer.Addr))

	err := s.httpServer.ListenAndServe()
	if err != nil {
		s.logger.Error("Server failed to start", zap.Error(err))
	}
}

func ginZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logger.Info("Request handler",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("duration", time.Since(start)),
			zap.Int("response_size", c.Writer.Size()),
		)
	}
}
