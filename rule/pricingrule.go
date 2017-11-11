package rule

import "github.com/gauravthadani/checkout/product"

type Rulable interface {
	Evaluate(cart []product.Product) float64
}

type PricingRules []Rulable
