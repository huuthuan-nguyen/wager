package transformer

import "github.com/huuthuan-nguyen/wager/app/model"

type PurchaseTransformer struct {
	ID          int     `json:"id"`
	WagerID     int     `json:"wager_id"`
	BuyingPrice float64 `json:"buying_price"`
	BoughtAt    int64   `json:"bought_at"`
}

// Transform /**
func (purchase PurchaseTransformer) Transform(e interface{}) interface{} {
	p, ok := e.(model.Purchase)
	if !ok {
		return e
	}

	purchase.ID = p.ID
	purchase.WagerID = p.WagerID
	purchase.BuyingPrice = p.BuyingPrice
	purchase.BoughtAt = p.BoughtAt.Unix()

	return purchase
}
