package router

import (
	"booking-system/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()

	// Global/Public Endpoints
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"status": "open for business",
		})
	})

	return r
}
