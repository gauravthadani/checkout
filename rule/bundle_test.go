package rule

import (
	"github.com/gauravthadani/checkout/product"
	"testing"
)

func TestBundleRule(t *testing.T) {
	productDef := NewSampleProduct()
	bundleDef := NewBundleProduct()
	rule := NewBundleRule(productDef, bundleDef)

	cart := []product.Product{
		NewSampleProduct(),
		NewBundleProduct(),
	}

	discount := rule.Evaluate(cart)
	if discount != -5.0 {
		t.Errorf("expected %f, got %f as discount for the bundle product", -5.0, discount)
	}
}

func TestBundleRule_MoreProductsThanBundled(t *testing.T) {
	productDef := NewSampleProduct()
	bundleDef := NewBundleProduct()
	rule := NewBundleRule(productDef, bundleDef)

	cart := []product.Product{
		NewSampleProduct(),
		NewSampleProduct(),
		NewBundleProduct(),
	}

	discount := rule.Evaluate(cart)
	if discount != -5.0 {
		t.Errorf("expected %f, got %f as discount for the bundle products", -5.0, discount)
	}
}

func NewBundleProduct() product.Product {
	return product.Product{
		SKU:   "cde",
		Name:  "sample cable",
		Price: 5.0,
	}
}
