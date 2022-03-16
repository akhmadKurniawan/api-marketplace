package models

import "time"

type Model struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
