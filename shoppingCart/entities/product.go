package entities


type Product struct {
	Id int64 'json: "id"'
	Name string 'json: "Name"'
	Price float64 'json: "Price"'
	Quantity int64 'json: "Quantity"'
	Photo string  'json: "Photo"'
}