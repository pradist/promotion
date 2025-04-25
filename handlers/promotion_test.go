package handlers_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/pradist/promotion/handlers"
	"github.com/pradist/promotion/services/mock_services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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

		gin.SetMode(gin.TestMode)
		res := httptest.NewRecorder()
		ctx, r := gin.CreateTestContext(res)
		r.GET("/calculate", promoHandler.CalculateDiscount)
		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		if err != nil {
			t.Errorf("got error: %s", err)
		}

		//Act
		r.ServeHTTP(res, req)

		//Assert
		assert.Equal(t, http.StatusOK, res.Code)
		body, _ := io.ReadAll(res.Body)
		assert.Equal(t, strconv.Itoa(expected), string(body))
	})
	t.Run("failure_badRequest", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		service := mock_services.NewMockPromotionService(ctrl)
		promoHandler := handlers.NewPromotionHandler(service)

		gin.SetMode(gin.TestMode)
		res := httptest.NewRecorder()
		ctx, r := gin.CreateTestContext(res)
		r.GET("/calculate", promoHandler.CalculateDiscount)

		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("/calculate?amount=%v", "A"), nil)
		if err != nil {
			t.Errorf("got error: %s", err)
		}

		//Act
		r.ServeHTTP(res, req)

		//Assert
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("failure_serviceReturnError", func(t *testing.T) {
		amount := 100

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		service := mock_services.NewMockPromotionService(ctrl)
		service.EXPECT().CalculateDiscount(amount).Return(0, errors.New("error"))
		promoHandler := handlers.NewPromotionHandler(service)

		gin.SetMode(gin.TestMode)
		res := httptest.NewRecorder()
		ctx, r := gin.CreateTestContext(res)
		r.GET("/calculate", promoHandler.CalculateDiscount)

		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		if err != nil {
			t.Errorf("got error: %s", err)
		}

		//Act
		r.ServeHTTP(res, req)

		//Assert
		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
