package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID `gorm:"primaryKey"`
	lineId     string
	oauthToken string
	createdAt  time.Time
	updatedAt  time.Time
}
