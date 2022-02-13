package request

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Wager struct {
	TotalWagerValue   int     `json:"total_wager_value" validate:"required"`
	Odds              int     `json:"odds" validate:"required"`
	SellingPercentage int     `json:"selling_percentage" validate:"required"`
	SellingPrice      float64 `json:"selling_price" validate:"required"`
}

func (wager *Wager) Validate() error {
	validate := validator.New()
	return validate.Struct(wager)
}

func (wager *Wager) Bind(r *http.Request) error {
	return nil
}
