package models

type Seller struct {
	Model
	UserID User
	Name   string `json:"name"`
	Almat  string `json:"alamt"`
	NoHp   int    `json:"no_hp"`
}
