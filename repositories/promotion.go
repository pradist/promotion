package repositories

//go:generate mockgen -source=./promotion.go -destination=mock_repositories/promotion.go

type PromotionRepository interface {
	GetPromotion() (Promotion, error)
}

type Promotion struct {
	ID              int
	PurchaseMin     int
	DiscountPercent int
}

type promotionRepository struct {
}

func NewPromotionRepository() PromotionRepository {
	return &promotionRepository{}
}

func (r *promotionRepository) GetPromotion() (Promotion, error) {
	return Promotion{
		ID:              1,
		PurchaseMin:     100,
		DiscountPercent: 10,
	}, nil
}
