package models

type UserToken struct {
	Model
	UserID int    `json:"user_id" db:"user_id"`
	Token  string `json:"token" db:"token"`
}
