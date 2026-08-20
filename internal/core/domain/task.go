package domain

import (
	"fmt"
	"time"
	core_error "todolist/internal/core/errors"

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

func NewTask(
	id uuid.UUID,
	version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorUserID uuid.UUID,
) Task {
	return Task{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		Completed:    completed,
		CreatedAt:    createdAt,
		CompletedAt:  completedAt,
		AuthorUserID: authorUserID,
	}
}

func CreateTask(
	title string,
	description *string,
	authorUserID uuid.UUID,
) Task {
	var (
		id                     = uuid.New()
		version                = 1
		completed              = false
		createdAt              = time.Now()
		completedAt *time.Time = nil
	)

	return NewTask(
		id,
		version,
		title,
		description,
		completed,
		createdAt,
		completedAt,
		authorUserID,
	)
}

func (t *Task) Validate() error {
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"invalid `Title` len: %d: %w",
			titleLen,
			core_error.ErrInvalidArgument,
		)
	}

	if t.Description != nil {
		descriptionLen := len([]rune(*t.Description))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf(
				"invalid `Description` len: %d: %w",
				descriptionLen,
				core_error.ErrInvalidArgument,
			)
		}
	}

	// Инвариант: Completed и CompletedAt должны быть согласованы.
	// Если задача выполнена, CompletedAt обязателен и не может быть раньше CreatedAt.
	// Если задача не выполнена, CompletedAt должен быть nil.
	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf(
				"`CompletedAt` can't be `nil` if `Completed`==`true`: %w",
				core_error.ErrInvalidArgument,
			)
		}

		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf(
				"`CompletedAt` can't be before `CreatedAt`: %w",
				core_error.ErrInvalidArgument,
			)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf(
				"`CompletedAt` must be `nil` if `Completed`==`false`: %w",
				core_error.ErrInvalidArgument,
			)
		}
	}

	return nil
}
