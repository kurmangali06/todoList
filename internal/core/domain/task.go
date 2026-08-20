package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID      uuid.UUID
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserID uuid.UUID
}
