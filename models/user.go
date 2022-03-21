package models

import "time"

type User struct {
	Model
	Email       string      `json:"email" db:"email"`
	Username    string      `json:"username" db:"username"`
	Password    string      `json:"password" db:"password"`
	Role        int         `json:"role" db:"role"`
	Status      string      `json:"status" db:"status"`
	LastLoginAt time.Time   `json:"last_login_at" db:"last_login_at"`
	UserToken   UserToken   `json:"token"`
	Transaction Transaction `json:"transaction"`
}

type UserParams struct {
	TransactionID int `json:"transaction_id"`
}
