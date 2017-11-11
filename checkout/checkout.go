package checkout

import (
	"github.com/gauravthadani/checkout/product"
	"github.com/gauravthadani/checkout/rule"
)

type Biller interface {
	Scan(p product.Product) error
	Total() (float64, error)
}

type StoreCheckout struct {
	cart []product.Product
}

func NewStoreCheckout(rules rule.PricingRules) *StoreCheckout {
	return &StoreCheckout{
		cart:[]product.Product{},
	}
}

func (sc *StoreCheckout) Scan(p product.Product) error {
	sc.cart = append(sc.cart, p)
	return nil
}

func (sc *StoreCheckout) Total() (float64, error) {
	total := 0.0
	for _, item := range sc.cart {
		total += item.Price
	}
	return total, nil
}
