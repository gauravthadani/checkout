package rule

import (
	"testing"

	"github.com/gauravthadani/checkout/product"
)

func TestFixedPrice(t *testing.T) {
	productType := NewSampleProduct()
	rule := NewFixedPriceRule(productType.SKU, 3, 50.0)

	cart := []product.Product{
		NewSampleProduct(),
		NewSampleProduct(),
		NewSampleProduct(),
	}

	discount := rule.Evaluate(cart)
	if discount != -150.0 {
		t.Errorf("expected %f, got %f as discount for bulk quantity", 150.0, discount)
	}
}
