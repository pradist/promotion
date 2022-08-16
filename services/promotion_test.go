package services_test

import (
	"errors"
	"gotest/repositories"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	testCases := []testCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "not_apply_80", purchaseMin: 100, discountPercent: 20, amount: 80, expected: 80},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			//Arrage
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)
			promoSerivce := services.NewPromotionService(promoRepo)

			// Act
			discount, _ := promoSerivce.CalculateDiscount(c.amount)
			expected := c.expected

			//Assert
			assert.Equal(t, expected, discount)
		})
	}

	t.Run("TestErrorZeroAmount", func(t *testing.T) {
		//Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 50,
		}, nil)
		promoSerivce := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoSerivce.CalculateDiscount(0)

		//Assert
		assert.Error(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("Purchase amount zero", func(t *testing.T) {
		//Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New("Database error"))
		promoSerivce := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoSerivce.CalculateDiscount(50)

		//Assert
		assert.Error(t, err, services.ErrRepository)
	})
}
