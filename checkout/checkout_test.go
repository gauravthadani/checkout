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

func TestCheckoutTotal_SampleScenario(t *testing.T) {

	appleTVDef := NewAppleTV()
	pricingRules := rule.PricingRules{
		rule.NewThreeForTwoPricingRule(appleTVDef),
	}

	checkout := NewStoreCheckout(pricingRules)

	checkout.Scan(NewAppleTV())
	checkout.Scan(NewAppleTV())
	checkout.Scan(NewAppleTV())
	checkout.Scan(NewVGA())

	total, _ := checkout.Total()
	if (total != 249.0) {
		t.Errorf("Expected $%f, got $%f for 3 Apple TV's and a VGA adapter", 249.0, total)
	}
}

func TestCheckoutTotal_SampleScenario2(t *testing.T) {

	ipadDef := NewIPad()
	checkout := NewStoreCheckout(rule.PricingRules{
		rule.NewBulkPriceRule(ipadDef, 4, 499.99),
	})

	checkout.Scan(NewAppleTV())
	checkout.Scan(NewIPad())
	checkout.Scan(NewIPad())
	checkout.Scan(NewAppleTV())
	checkout.Scan(NewIPad())
	checkout.Scan(NewIPad())
	checkout.Scan(NewIPad())

	total, _ := checkout.Total()
	if (total != 2718.95) {
		t.Errorf("Expected $%f, got $%f for 5 Apple TV's and 2 IPads", 2718.95, total)
	}
}
//
//func TestCheckoutTotal_SampleScenario3(t *testing.T){
//
//	mbpDef := NewMacBookPro()
//	vgaDef:= NewVGA()
//
//	checkout := NewStoreCheckout(nil)
//
//	checkout.Scan(NewMacBookPro())
//	checkout.Scan(NewVGA())
//	checkout.Scan(NewIPad())
//
//	total, _ := checkout.Total()
//	if (total != 1949.98) {
//		t.Errorf("Expected $%f, got $%f for 1 MacBookPro, 1 VGA and 1 Ipad", 1949.98, total)
//	}
//}


func NewVGA() product.Product {
	return product.Product{
		Name:"VGA adapter",
		SKU:"vga",
		Price:30.0,
	}
}

func NewIPad() product.Product {
	return product.Product{
		Name:"Super iPad",
		SKU:"ipd",
		Price:549.99,
	}
}

func NewAppleTV() product.Product {
	return product.Product{
		Name:"Apple TV",
		SKU:"atv",
		Price:109.50,
	}
}

func NewMacBookPro() product.Product {
	return product.Product{
		Name:"MacBook Pro",
		SKU:"mbp",
		Price:1399.99,
	}
}



