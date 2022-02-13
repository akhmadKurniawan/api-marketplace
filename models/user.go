package models

import "time"

type User struct {
	Model
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password"`
	Role        int       `json:"role" db:"role"`
	LastLoginAt time.Time `json:"last_login_at" db:"last_login_at"`
	UserToken   UserToken `json:"token"`
}
