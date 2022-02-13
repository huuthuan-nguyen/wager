package transformer

import (
	"github.com/huuthuan-nguyen/wager/app/model"
)

type WagerTransformer struct {
	ID                  int      `json:"id"`
	TotalWagerValue     int      `json:"total_wager_value"`
	Odds                int      `json:"odds"`
	SellingPercentage   int      `json:"selling_percentage"`
	SellingPrice        float64  `json:"selling_price"`
	CurrentSellingPrice float64  `json:"current_selling_price"`
	PercentageSold      *float64 `json:"percentage_sold"`
	AmountSold          *float64 `json:"amount_sold"`
	PlacedAt            int64    `json:"placed_at"`
}

// Transform /**
func (wager WagerTransformer) Transform(e interface{}) interface{} {
	w, ok := e.(model.Wager)
	if !ok {
		return e
	}

	wager.ID = w.ID
	wager.TotalWagerValue = w.TotalWagerValue
	wager.Odds = w.Odds
	wager.SellingPercentage = w.SellingPercentage
	wager.SellingPrice = w.SellingPrice
	wager.CurrentSellingPrice = w.CurrentSellingPrice
	wager.PercentageSold = w.PercentageSold
	wager.AmountSold = w.AmountSold
	wager.PlacedAt = w.PlacedAt.Unix()

	return wager
}
