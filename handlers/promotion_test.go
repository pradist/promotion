package handlers_test

import (
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/pradist/promotion/handlers"
	"github.com/pradist/promotion/services/mock_services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		service := mock_services.NewMockPromotionService(ctrl)
		service.EXPECT().CalculateDiscount(amount).Return(expected, nil)
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
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		service := mock_services.NewMockPromotionService(ctrl)
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

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		service := mock_services.NewMockPromotionService(ctrl)
		service.EXPECT().CalculateDiscount(amount).Return(0, errors.New("error"))
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
