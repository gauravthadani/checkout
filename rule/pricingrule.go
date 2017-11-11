package rule

import "github.com/gauravthadani/checkout/product"

type Rulable interface {
	Evaluate(items []product.Product) (float64 ,error)
}

type PricingRules []Rulable