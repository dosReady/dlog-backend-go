package models

import (
	"time"
)

type Base struct {
	CreatedID string
	UpdatedID string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
