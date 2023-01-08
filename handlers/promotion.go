package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/pradist/promotion/services"
)

type PromotionHandler interface {
	CalculateDiscount(c *gin.Context)
}

type promotionHandler struct {
	promoService services.PromotionService
}

func NewPromotionHandler(promoService services.PromotionService) PromotionHandler {
	return promotionHandler{promoService: promoService}
}

func (h promotionHandler) CalculateDiscount(c *gin.Context) {
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discount, err := h.promoService.CalculateDiscount(amount)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "%d", discount)
}
