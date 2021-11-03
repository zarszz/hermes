package entity

type ProductInput struct {
	Id int `json:"id"`
	Sku  string `json:"sku"`
	Name string `json:"name"`
	Display string `json:"display"`
}

