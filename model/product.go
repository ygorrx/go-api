package model

type Product struct {
	ID    int   `json:"id_product"`
	Name  string  `json:"name_product"`
	Price float64  `json:"price_product"`
}