package rule

import "github.com/gauravthadani/checkout/product"

type ThreeForTwoPricingRule struct {

}

func NewThreeForTwoPricingRule() *ThreeForTwoPricingRule {
	return &ThreeForTwoPricingRule{
	}
}

func (rule *ThreeForTwoPricingRule) Evaluate(items []product.Product) (float64, error) {
	return -100, nil
}



