package models

import (
	"time"
)

type ItemsReturn struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	RequestID  int       `json:"request_id" gorm:"not null"` // Foreign key to ItemsRequest
	ReturnDate time.Time `json:"return_date" gorm:"autoCreateTime"`
}
