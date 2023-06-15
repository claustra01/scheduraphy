package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id           uuid.UUID `gorm:"primaryKey"`
	lineId       string
	refreshToken string
	createdAt    time.Time
	updatedAt    time.Time
}
