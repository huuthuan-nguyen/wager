package request

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Purchase struct {
	BuyingPrice float64 `json:"buying_price" validate:"required"`
}

func (purchase *Purchase) Validate() error {
	validate := validator.New()
	return validate.Struct(purchase)
}

func (purchase *Purchase) Bind(r *http.Request) error {
	return nil
}
