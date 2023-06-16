package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `gorm:"primaryKey"`
	LineId       string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
