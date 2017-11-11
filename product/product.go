package product

type Product struct {
	SKU   string
	Name  string
	Price float64
}

type Catalog []Product
