package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pradist/promotion/handlers"
	"github.com/pradist/promotion/repositories"
	"github.com/pradist/promotion/services"
)

var setTrustedProxies = func(app *gin.Engine) error {
	return app.SetTrustedProxies(nil)
}

func New() *gin.Engine {
	r := repositories.NewPromotionRepository()
	s := services.NewPromotionService(r)
	h := handlers.NewPromotionHandler(s)

	app := gin.Default()
	app.GET("/health", Health)
	app.GET("/version", Version)
	app.GET("/api/v1/promotion", h.CalculateDiscount)
	if err := setTrustedProxies(app); err != nil {
		panic(err)
	}

	return app
}
