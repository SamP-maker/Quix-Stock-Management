package models

import (
	"time"
)

type ItemsStatus struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ReturnID  int       `json:"return_id" gorm:"not null"` // Foreign key to ItemsReturn
	ItemID    int       `json:"item_id" gorm:"not null"`
	Status    string    `json:"status" validate:"required" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
