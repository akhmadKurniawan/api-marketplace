package models

type ProductType struct {
	Model
	Name  string `json:"name" db:"name"`
	Image string `json:"image" db:"image"`
}
