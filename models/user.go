package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;uniqueIndex"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
