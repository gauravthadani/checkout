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
	cart         []product.Product
	pricingRules rule.PricingRules
}

func NewStoreCheckout(rules rule.PricingRules) *StoreCheckout {
	return &StoreCheckout{
		cart:         []product.Product{},
		pricingRules: rules,
	}
}

func (sc *StoreCheckout) Scan(p product.Product) error {
	return sc.scan(p)
}

func (sc *StoreCheckout) scan(p ...product.Product) error {
	sc.cart = append(sc.cart, p ...)
	return nil
}

func (sc *StoreCheckout) Total() (float64, error) {
	total := 0.0
	for _, item := range sc.cart {
		total += item.Price
	}
	for _, rule := range sc.pricingRules {
		total += rule.Evaluate(sc.cart)
	}
	return total, nil
}
