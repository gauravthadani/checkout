package checkout

import (
	"github.com/gauravthadani/checkout/product"
	"github.com/gauravthadani/checkout/rule"
)

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

func (sc *StoreCheckout) Scan(p product.Product) {
	sc.scan(p)
}

func (sc *StoreCheckout) scan(p ...product.Product) {
	sc.cart = append(sc.cart, p...)
}

func (sc *StoreCheckout) Total() float64 {
	total := 0.0
	for _, item := range sc.cart {
		total += item.Price
	}
	for _, rule := range sc.pricingRules {
		total += rule.Evaluate(sc.cart)
	}
	return total
}
