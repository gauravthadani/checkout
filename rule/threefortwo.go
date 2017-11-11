package rule

import "github.com/gauravthadani/checkout/product"

type ThreeForTwoPricingRule struct {
	applicationSKU string
	price          float64
}

func NewThreeForTwoPricingRule(sku string, price float64) *ThreeForTwoPricingRule {
	return &ThreeForTwoPricingRule{
		applicationSKU:sku,
		price:price,
	}
}

func (rule *ThreeForTwoPricingRule) Evaluate(cart []product.Product) float64 {
	count := 0
	for _, item := range cart {
		if (item.SKU == rule.applicationSKU) {
			count++
		}
	}

	if (count >= 3) {
		return -rule.price
	}
	return 0
}



