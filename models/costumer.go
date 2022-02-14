package models

type Costumer struct {
	Model
	UserID int    `json:"user_id" db:"user_id"`
	Name   string `json:"name" db:"name"`
	Alamat string `json:"alamat" db:"alamat"`
	NoHp   string `json:"no_hp" db:"no_hp"`
}
