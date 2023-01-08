package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pradist/promotion/handlers"
	"github.com/pradist/promotion/repositories"
	"github.com/pradist/promotion/services"
)

func New() *gin.Engine {
	r := repositories.NewPromotionRepository()
	s := services.NewPromotionService(r)
	h := handlers.NewPromotionHandler(s)

	app := gin.Default()
	app.GET("/health", Health)
	app.GET("/api/v1/promotion", h.CalculateDiscount)

	return app
}
