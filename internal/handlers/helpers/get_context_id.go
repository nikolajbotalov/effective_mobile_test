package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetContextID(c *gin.Context, logger *zap.Logger) (string, error) {
	logger.Info("Getting id from context")

	id, exists := c.Get("id")
	if !exists {
		logger.Error("id is required")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id is required"})
		return "", errors.New("id is required")
	}

	idStr, ok := id.(string)
	if !ok {
		logger.Error("invalid id type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid id type"})
		return "", errors.New("invalid id type")
	}

	return idStr, nil
}
