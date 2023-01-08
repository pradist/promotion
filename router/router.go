package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pradist/promotion/handlers"
	"github.com/pradist/promotion/repositories"
	"github.com/pradist/promotion/services"
	"net/http"
)

func New() *gin.Engine {
	r := repositories.NewPromotionRepository()
	s := services.NewPromotionService(r)
	h := handlers.NewPromotionHandler(s)

	app := gin.Default()
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	app.GET("/api/v1/promotion", h.CalculateDiscount)
	return app
}
