package model

type Product struct {
	ID           int     `json:"id_product"`
	Product_Name string  `json:"name"`
	Price        float64 `json:"price"`
}
