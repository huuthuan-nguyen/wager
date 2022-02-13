package model

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Purchase struct {
	bun.BaseModel `bun:"table:purchases,alias:purchase"`
	ID            int       `bun:"id,pk,autoincrement"`
	WagerID       int       `bun:"wager_id"`
	BuyingPrice   float64   `bun:"buying_price"`
	BoughtAt      time.Time `bun:"bought_at"`
	Wager         *Wager    `bun:"rel:belongs-to,join=wager_id=id"`
}

// Create new Purchase record or Buy Wager
func (purchase *Purchase) Create(ctx context.Context, db *bun.DB) error {
	_, err := db.NewInsert().Model(purchase).Exec(ctx)
	return err
}
