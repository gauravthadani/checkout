package rule

import (
	"testing"

	"github.com/gauravthadani/checkout/product"
)

func Test3for2PricingRule(t *testing.T) {
	rule := NewThreeForTwoPricingRule()

	cartItems := []product.Product{NewSampleProduct(), NewSampleProduct(), NewSampleProduct() }

	discount, _ := rule.Evaluate(cartItems)
	if (discount != -100.0) {
		t.Errorf("expected %f, got %f to discount 3 items for the price of 2", -100, discount)
	}

}

func NewSampleProduct() product.Product {
	return product.Product{
		Name:"MyItem",
		SKU:"abc",
		Price:100,
	}
}
