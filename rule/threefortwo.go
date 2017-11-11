package rule

import "github.com/gauravthadani/checkout/product"

type ThreeForTwoPricingRule struct {
	sku string
}

func NewThreeForTwoPricingRule(sku string) *ThreeForTwoPricingRule {
	return &ThreeForTwoPricingRule{
		sku:sku,
	}
}

func (rule *ThreeForTwoPricingRule) Evaluate(cart []product.Product) (float64, error) {
	count := 0

	for _, item := range cart {
		if (item.SKU == rule.sku) {
			count++
		}
	}

	if (count >= 3) {
		return -100, nil
	}
	return 0, nil
}



