package checkout

import (
	"testing"

	"github.com/gauravthadani/checkout/product"
	"github.com/gauravthadani/checkout/rule"
)

func TestCheckoutTotalAppleTVs(t *testing.T) {
	checkout := NewStoreCheckout(rule.PricingRules{})

	checkout.Scan(NewAppleTV())
	checkout.Scan(NewAppleTV())

	total, _ := checkout.Total()
	if (total != 219.0) {
		t.Errorf("Expected $%f, got $%f for 2 Apple TV's", 219.0, total)
	}
}

func TestCheckoutTotalSampleScenario(t *testing.T){

	//pricingRules := rule.PricingRules{
	//	rule.PricingRule{
	//
	//	},
	//}

	checkout := NewStoreCheckout(rule.PricingRules{})

	checkout.Scan(NewAppleTV())
	checkout.Scan(NewAppleTV())
	//checkout.Scan(NewAppleTV())

	checkout.Scan(NewVGA())

	total, _:=checkout.Total()
	if(total != 249.0) {
		t.Errorf("Expected $%f, got $%f for 3 Apple TV's and a VGA adapter", 249.0, total)
	}
}

func NewVGA() product.Product{
	return product.Product{
		Name:"VGA adapter",
		SKU:"vga",
		Price:30.0,
	}
}

func NewAppleTV() product.Product {
	return product.Product{
		Name:"Apple TV",
		SKU:"atv",
		Price:109.50,
	}
}


