package models

import (
	"time"

	"gorm.io/gorm"
)

type Agenda struct {
	ID               uint   `gorm:"primaryKey"`
	UserId           uint   `gorm:"not null;index"`
	Title            string `gorm:"size:255;not null"`
	Category         string `gorm:"size:100;not null"`
	Description      string `gorm:"type:text"`
	Status           uint   `gorm:"default:0"`
	IsLawyerAssigned bool   `gorm:"default:false"`
	LawyerId         uint   `gorm:"index"`
	MeetingDate      *time.Time
	IsEvaluationDone bool `gorm:"default:false"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
