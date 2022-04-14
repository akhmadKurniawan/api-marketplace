package models

import (
	"time"
)

type User struct {
	ID          int         `gorm:"primary_key" json:"id"`
	Email       string      `json:"email" db:"email"`
	Username    string      `json:"username" db:"username"`
	Password    string      `json:"password" db:"password"`
	Role        int         `json:"role" db:"role"`
	Status      string      `json:"status" db:"status"`
	LastLoginAt time.Time   `json:"last_login_at" db:"last_login_at"`
	UserToken   UserToken   `json:"token"`
	Transaction Transaction `json:"transaction"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type UserParams struct {
	TransactionID int `json:"transaction_id"`
}
