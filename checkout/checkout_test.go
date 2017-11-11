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

	checkout := NewCheckout()
	checkout.Scan(atv1)
	checkout.Scan(atv2)

	if (checkout.Total() != 219.0){
		t.Errorf("Expected %f, got %f",219.0, checkout.Total())
	}

}
