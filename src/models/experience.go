package model

import (
	"time"
)

type Experience struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Company     string    `json:"company" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Experience) TableName() string {
	return "experiences"
}
