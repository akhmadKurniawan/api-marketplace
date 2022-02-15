package models

type Transaction struct {
	Model
	UserID       int    `json:"user_id" db:"user_id"`
	ProductID    int    `json:"product_id" db:"product_id"`
	Type         string `json:"type" db:"type"`
	Description  string `json:"description" db:"description"`
	Amount       int    `json:"amount" db:"amount"`
	TotalProduct string `json:"total_product" db:"total_product"`
}
