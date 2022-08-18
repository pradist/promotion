package handlers_test

import (
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pradist/promotion/handlers"
	"github.com/stretchr/testify/assert"

	serviceMock "github.com/pradist/promotion/mocks/services"
)

func TestCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		service := &serviceMock.PromotionService{}
		service.On("CalculateDiscount", amount).Return(expected, nil)
		promoHandler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})
	t.Run("failure_badRequest", func(t *testing.T) {
		service := &serviceMock.PromotionService{}
		promoHandler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", "A"), nil)

		//Act
		res, _ := app.Test(req)
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		//Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})
	t.Run("failure_serviceReturnError", func(t *testing.T) {
		amount := 100

		service := &serviceMock.PromotionService{}
		service.On("CalculateDiscount", amount).Return(0, errors.New("error"))
		promoHandler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		//Assert
		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
	})
}
