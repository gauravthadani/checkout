package factory

import (
	"github.com/gauravthadani/checkout/rule"
	"github.com/gauravthadani/checkout/product"
)

type factory struct {
	catalog product.Catalog
	rules   rule.PricingRules
}

func (factory factory) PricingRules() rule.PricingRules {
	return factory.rules
}

func NewFactory() *factory {
	vga := product.Product{
		Name:  "VGA adapter",
		SKU:   "vga",
		Price: 30.0,
	}

	ipad := product.Product{
		Name:  "Super iPad",
		SKU:   "ipd",
		Price: 549.99,
	}

	appleTV := product.Product{
		Name:  "Apple TV",
		SKU:   "atv",
		Price: 109.50,
	}

	macBookPro := product.Product{
		Name:  "MacBook Pro",
		SKU:   "mbp",
		Price: 1399.99,
	}

	catalog := []product.Product{vga, ipad, appleTV, macBookPro}

	rules := rule.PricingRules{
		rule.NewThreeForTwoPricingRule(appleTV),
		rule.NewBulkPriceRule(ipad, 4, 499.99),
		rule.NewBundleRule(macBookPro, vga),
	}

	return &factory{
		catalog:catalog,
		rules:rules,
	}
}


