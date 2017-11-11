package rule

import "github.com/gauravthadani/checkout/product"

type bundleRule struct {
	applicableProduct product.Product
	bundleProduct     product.Product
}

func NewBundleRule(product product.Product, bundleProduct product.Product) *bundleRule {
	return &bundleRule{
		applicableProduct:product,
		bundleProduct:bundleProduct,
	}
}

func (rule *bundleRule) Evaluate(cart []product.Product) float64 {
	productCount := 0
	for _, item := range cart {
		if item.SKU == rule.applicableProduct.SKU {
			productCount++
		}
	}

	bundledItemsCount:=0
	for _, item := range cart {
		if item.SKU == rule.bundleProduct.SKU {
			bundledItemsCount++
		}
	}


	return - (float64(productCount-bundledItemsCount) * rule.bundleProduct.Price)

}