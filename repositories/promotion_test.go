package repositories

import "testing"

func TestGetPromotion(t *testing.T) {
	repo := NewPromotionRepository()
	promotion, err := repo.GetPromotion()
	if err != nil {
		t.Errorf("GetPromotion() error = %v", err)
		return
	}
	if promotion.ID != 1 {
		t.Errorf("GetPromotion() promotion.ID = %v, want %v", promotion.ID, 1)
	}
	if promotion.PurchaseMin != 100 {
		t.Errorf("GetPromotion() promotion.PurchaseMin = %v, want %v", promotion.PurchaseMin, 100)
	}
	if promotion.DiscountPercent != 10 {
		t.Errorf("GetPromotion() promotion.DiscountPercent = %v, want %v", promotion.DiscountPercent, 10)
	}
}
