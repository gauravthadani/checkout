package checkout

import "github.com/gauravthadani/checkout/product"

type Biller interface {
	Scan(p product.Product) error
	Total() (float64, error)
}

type StoreCheckout struct {

}

func NewStoreCheckout() *StoreCheckout {
	return &StoreCheckout{}

}

func (sc *StoreCheckout) Scan(p product.Product) error {
	return nil
}

func (sc *StoreCheckout) Total() (float64, error) {
	return 219.0, nil
}
