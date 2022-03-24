package models

type ProductType struct {
	Model
	Name  string `bson:"name" json:"name" db:"name"`
	Image string `bson:"image" json:"image" db:"image"`
}
