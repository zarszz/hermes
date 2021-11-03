package models

type Product struct {
	Id int
	Sku  string
	Name string
	Display string
}

type Products struct {
	Products []Product
}


