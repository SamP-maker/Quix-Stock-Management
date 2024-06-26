package models

import (
	"time"
)

type ItemsRequest struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ItemID       string    `json:"item_id" validate:"required,min=2,max=100" gorm:"unique;not null"`
	ItemName     string    `json:"item_name" validate:"required,min=2,max=100" gorm:"unique;not null"`
	ItemQuantity int       `json:"item_quantity" gorm:"not null"`
	RequestDate  time.Time `json:"request_date" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
