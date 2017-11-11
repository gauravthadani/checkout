package rule

import "github.com/gauravthadani/checkout/product"

type threeForTwoPricingRule struct {
	applicableProduct product.Product
}

func NewThreeForTwoPricingRule(p product.Product) *threeForTwoPricingRule {
	return &threeForTwoPricingRule{
		applicableProduct: p,
	}
}

func (rule *threeForTwoPricingRule) Evaluate(cart []product.Product) float64 {
	count := 0
	for _, item := range cart {
		if item.SKU == rule.applicableProduct.SKU {
			count++
		}
	}

	if count >= 3 {
		return -rule.applicableProduct.Price
	}
	return 0
}
