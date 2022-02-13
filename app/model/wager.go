package model

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Wager struct {
	bun.BaseModel       `bun:"table:purchases,alias:purchase"`
	ID                  int         `bun:"id,pk,autoincrement"`
	TotalWagerValue     int         `bun:"total_wager_value"`
	Odds                int         `bun:"odds"`
	SellingPercentage   int         `bun:"selling_percentage"`
	SellingPrice        float64     `bun:"selling_price"`
	CurrentSellingPrice float64     `bun:"current_selling_price"`
	PercentageSold      *float64    `bun:"percentage_sold"`
	AmountSold          *float64    `bun:"amount_sold"`
	PlacedAt            time.Time   `bun:"placed_at"`
	Purchases           []*Purchase `bun:"rel:has-many,join:id=wager_id"`
}

// Create new wager or Place Wager
func (wager *Wager) Create(ctx context.Context, db *bun.DB) error {
	_, err := db.NewInsert().Model(wager).Exec(ctx)
	return err
}

// FindAllWagers Fetch all wagers
func FindAllWagers(ctx context.Context, db *bun.DB) ([]Wager, error) {
	wagers := make([]Wager, 0)
	err := db.NewSelect().Model(&wagers).Scan(ctx)
	return wagers, err
}
