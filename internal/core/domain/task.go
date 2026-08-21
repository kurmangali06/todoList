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

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskPatch(
	title Nullable[string],
	description Nullable[string],
	completed Nullable[bool],
) TaskPatch {
	return TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (t *TaskPatch) Validate() error {
	if t.Title.Set && t.Title.Value == nil {
		return fmt.Errorf("`Title` is required: %w", core_error.ErrInvalidArgument)
	}
	if t.Completed.Set && t.Completed.Value == nil {
		return fmt.Errorf("`Completed` is required: %w", core_error.ErrInvalidArgument)
	}
	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate task patch: %w", err)
	}
	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}
	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Completed.Set {
		completed := *patch.Completed.Value

		if completed {
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt
		} else {
			tmp.CompletedAt = nil
		}

		tmp.Completed = completed
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched task: %w", err)
	}

	*t = tmp
	return nil
}
