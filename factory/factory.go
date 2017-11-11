package factory

import (
	"github.com/gauravthadani/checkout/product"
	"github.com/gauravthadani/checkout/rule"
)

type factory struct {
	catalog product.Catalog
	rules   rule.PricingRules
}

func (factory factory) PricingRules() rule.PricingRules {
	return factory.rules
}

var VGA product.Product = product.Product{
	Name:  "VGA adapter",
	SKU:   "vga",
	Price: 30.0,
}

var Ipad product.Product = product.Product{
	Name:  "Super iPad",
	SKU:   "ipd",
	Price: 549.99,
}

var AppleTV product.Product = product.Product{
	Name:  "Apple TV",
	SKU:   "atv",
	Price: 109.50,
}

var MacBookPro product.Product = product.Product{
	Name:  "MacBook Pro",
	SKU:   "mbp",
	Price: 1399.99,
}

func (factory factory) ProductCatalog() product.Catalog {
	return factory.catalog
}

func NewFactory() *factory {

	catalog := []product.Product{VGA, Ipad, AppleTV, MacBookPro}

	rules := rule.PricingRules{
		rule.NewThreeForTwoPricingRule(AppleTV),
		rule.NewBulkPriceRule(Ipad, 4, 499.99),
		rule.NewBundleRule(MacBookPro, VGA),
	}

	return &factory{
		catalog: catalog,
		rules:   rules,
	}
}
