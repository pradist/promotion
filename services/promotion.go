package services

import (
	"errors"

	"github.com/pradist/promotion/repositories"
)

//go:generate mockgen -source=./promotion.go -destination=mock_services/promotion.go

var (
	ErrZeroAmount = errors.New("purchase amount could not be zero")
	ErrRepository = errors.New("repository error")
)

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
	promoRepo repositories.PromotionRepository
}

func NewPromotionService(PromoRepo repositories.PromotionRepository) PromotionService {
	return promotionService{promoRepo: PromoRepo}
}

func (p promotionService) CalculateDiscount(amount int) (int, error) {
	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	promotion, err := p.promoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	if amount >= promotion.PurchaseMin {
		return amount - (promotion.DiscountPercent * amount / 100), nil
	}

	return amount, nil
}
