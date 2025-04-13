package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey"`
	Token     string     `gorm:"size:64;uniqueIndex"`
	UserID    uint       `gorm:"not null"`
	ExpiresAt time.Time  `gorm:"not null"`
	RevokedAt *time.Time `gorm:"default:null"`
	CreatedAt time.Time
}

func (r *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}

	if r.ExpiresAt.IsZero() {
		r.ExpiresAt = time.Now().Add(7 * 24 * time.Hour)
	}

	if r.Token == "" {
		r.Token = uuid.NewString()
	}

	return nil
}

func (r *RefreshToken) IsExpired() bool {
	return time.Now().After(r.ExpiresAt)
}

func (r *RefreshToken) IsRevoked() bool {
	return r.RevokedAt != nil
}
