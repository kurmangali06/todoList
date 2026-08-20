package tasks_postgres_repository

import (
	"time"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

type TaskModel struct {
	ID           uuid.UUID
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID uuid.UUID
}

func modelToDomain(model TaskModel) domain.Task {
	return domain.NewTask(
		model.ID,
		model.Version,
		model.Title,
		model.Description,
		model.Completed,
		model.CreatedAt,
		model.CompletedAt,
		model.AuthorUserID,
	)
}

func modelsToDomains(models []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(models))

	for i, model := range models {
		domains[i] = modelToDomain(model)
	}

	return domains
}
