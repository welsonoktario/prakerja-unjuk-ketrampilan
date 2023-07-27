package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"type:varchar(191);not null" json:"name"`
	Username  string         `gorm:"uniqueIndex;type:varchar(16);not null" json:"username"`
	Email     *string        `gorm:"uniqueIndex;type:varchar(191)" json:"email"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
