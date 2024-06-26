package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement`
	EID       string    `json:"username" validate:"required,min=2,max=100" gorm:"unique; not null"`
	Password  string    `json:"password" validate:"required,min=6" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"Updated_at" gorm:"autoUpdateTime"`
}
