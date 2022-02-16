package models

type Product struct {
	Model
	ProductType int    `json:"product_type" db:"product_type"`
	ShopId      int    `json:"shop_id" db:"shop_id"`
	Name        string `json:"name" db:"name"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
	Qty         string `json:"qty" db:"qty"`
	Image       string `json:"image" db:"image"`
	Shop        Shop   `json:"shop"`
}
