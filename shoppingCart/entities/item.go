package entities

type Item struct {
	Product Product 'json: "Product"'
	Quantity int64  'json: "Quantity"'
}
