package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	lineId     string
	googleId   string
	oauthToken string
	createdAt  time.Time
	updatedAt  time.Time
}
