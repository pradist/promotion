package repositories

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

func (p *promotionRepository) GetPromotion() (Promotion, error) {
	return Promotion{
		ID:              1,
		PurchaseMin:     100,
		DiscountPercent: 10,
	}, nil
}
