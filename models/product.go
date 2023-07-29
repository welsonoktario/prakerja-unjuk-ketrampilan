package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"type:varchar(191);not null" json:"name"`
	Price       uint           `gorm:"type:int unsigned;not null" json:"price"`
	Description *string        `gorm:"type:text" json:"description"`
	UserID      uint           `gorm:"not null" json:"-"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	User        User           `json:"user"`
}
