package models

import (
	"time"
)

type Inventory struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ItemName    string    `json:"item_name" gorm:"not null"` // Assuming this should be a string
	Quantity    int       `json:"quantity" validate:"required" gorm:"not null"`
	ItemInUse   int       `json:"item_in_use" validate:"required" gorm:"not null"`
	ItemBroken  int       `json:"item_broken" validate:"required" gorm:"not null"`
	ItemMissing int       `json:"item_missing" validate:"required" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
