package models

type Product struct {
	Model
	ShopId      int    `bson:"shop_id" json:"shop_id" db:"shop_id"`
	ProductType int    `bson:"product_type" json:"product_type" db:"product_type"`
	Name        string `bson:"name" json:"name" db:"name"`
	Price       int    `bson:"price" json:"price" db:"price"`
	Description string `bson:"description" json:"description" db:"description"`
	Qty         int    `bson:"qty" json:"qty" db:"qty"`
	Image       string `bson:"image" json:"image" db:"image"`
}
