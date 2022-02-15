package models

type Walet struct {
	Model
	UserID int `json:"user_id" db:"user_id"`
	Saldo  int `json:"saldo" db:"saldo"`
}
