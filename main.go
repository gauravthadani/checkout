package main

import (
	"fmt"

	"github.com/gauravthadani/checkout/checkout"
	"github.com/gauravthadani/checkout/factory"
)

func main() {

	f := factory.NewFactory()

	checkout := checkout.NewStoreCheckout(f.PricingRules())

	checkout.Scan(factory.VGA)
	checkout.Scan(factory.AppleTV)
	checkout.Scan(factory.MacBookPro)
	checkout.Scan(factory.Ipad)

	total := checkout.Total()
	fmt.Printf("The total bill amount is %f", total)

}
