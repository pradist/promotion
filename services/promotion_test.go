package services_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/pradist/promotion/repositories"
	"github.com/pradist/promotion/services"
	"github.com/stretchr/testify/assert"

	repoMock "github.com/pradist/promotion/mocks/repositories"
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
			//Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := repoMock.NewMockPromotionRepository(ctrl)
			repo.EXPECT().GetPromotion().Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)
			promoService := services.NewPromotionService(repo)

			// Act
			discount, _ := promoService.CalculateDiscount(c.amount)
			expected := c.expected

			//Assert
			assert.Equal(t, expected, discount)
		})
	}

	t.Run("TestErrorZeroAmount", func(t *testing.T) {
		//Arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := repoMock.NewMockPromotionRepository(ctrl)
		promoService := services.NewPromotionService(repo)

		// Act
		_, err := promoService.CalculateDiscount(0)

		//Assert
		assert.Error(t, err, services.ErrZeroAmount)
	})

	t.Run("Purchase amount zero", func(t *testing.T) {
		//Arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := repoMock.NewMockPromotionRepository(ctrl)
		repo.EXPECT().GetPromotion().Return(repositories.Promotion{}, errors.New("database error"))
		promoService := services.NewPromotionService(repo)

		// Act
		_, err := promoService.CalculateDiscount(50)

		//Assert
		assert.Error(t, err, services.ErrRepository)
	})
}
