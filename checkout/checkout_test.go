package checkout

import (
	"testing"

	"github.com/gauravthadani/checkout/product"
)

func TestCheckoutSimpleTotal(t *testing.T) {

	atv1 := product.Product{
		Name:"Apple TV",
		SKU:"atv",
		Price:109.50,
	}
	atv2 := product.Product{
		Name:"Apple TV",
		SKU:"atv",
		Price:109.50,
	}

	checkout := NewStoreCheckout()
	checkout.Scan(atv1)
	checkout.Scan(atv2)

	total, _ := checkout.Total()
	if (total != 219.0) {
		t.Errorf("Expected %f, got %f", 219.0, total)
	}

}
