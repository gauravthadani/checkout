package rule

import (
	"testing"

	"github.com/gauravthadani/checkout/product"
)

func TestFixedPrice(t *testing.T) {
	productdefinition := NewSampleProduct()
	rule := NewBulkPriceRule(productdefinition, 3, 40.0)

	cart := []product.Product{
		NewSampleProduct(),
		NewSampleProduct(),
		NewSampleProduct(),
	}

	discount := rule.Evaluate(cart)
	if discount != -180.0 {
		t.Errorf("expected %f, got %f as discount for bulk quantity", -180.0, discount)
	}
}
