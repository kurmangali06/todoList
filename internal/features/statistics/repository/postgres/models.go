package statistics_postgres_repository

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

func modelToDomain(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)
}

// modelsToDomains конвертирует список моделей в список доменных объектов.
func modelsToDomains(taskModels []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(taskModels))

	for i, model := range taskModels {
		domains[i] = modelToDomain(model)
	}

	return domains
}
