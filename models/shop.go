package models

type Shop struct {
	Model
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Alamat      string `json:"alamat" db:"alamat"`
	Logo        string `json:"logo" db:"logo"`
	SellerID    int    `json:"seller_id" db:"seller_id"`
	ProductID   int    `json:"product_id" db:"product_id"`
}
