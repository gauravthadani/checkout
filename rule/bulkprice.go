package rule

import "github.com/gauravthadani/checkout/product"

type bulkPriceRule struct {
	applicableProduct product.Product
	minQty            int
	pricePerItem      float64
}

func NewBulkPriceRule(product product.Product, bulkQty int, price float64) *bulkPriceRule {

	return &bulkPriceRule{
		applicableProduct: product,
		minQty:            bulkQty,
		pricePerItem:      price,
	}
}

func (rule *bulkPriceRule) Evaluate(cart []product.Product) float64 {
	count := 0
	for _, item := range cart {
		if item.SKU == rule.applicableProduct.SKU {
			count++
		}
	}

	if count > rule.minQty {
		discountPerItem := rule.applicableProduct.Price - rule.pricePerItem
		return -(discountPerItem * float64(count))
	}
	return 0

}
