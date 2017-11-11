package rule

import "github.com/gauravthadani/checkout/product"

type FixedPriceRule struct {
	applicableSKU string
	minQty        int
	pricePerItem  float64
}

func NewFixedPriceRule(sku string, bulkQty int, price float64) *FixedPriceRule {
	return &FixedPriceRule{
		applicableSKU: sku,
		minQty:   bulkQty,
		pricePerItem:  price,
	}
}

func (rule *FixedPriceRule) Evaluate(cart []product.Product) float64 {
	count := 0
	for _, item := range cart {
		if item.SKU == rule.applicableSKU {
			count++
		}
	}

	if count >= rule.minQty {
		return -(rule.pricePerItem * float64(count))
	}
	return 0

}
