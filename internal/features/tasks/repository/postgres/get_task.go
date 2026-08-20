package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"todolist/internal/core/domain"
	core_error "todolist/internal/core/errors"
	core_postgres_pool "todolist/internal/core/repository/postgres/pool"

	"github.com/google/uuid"
)

func (r *TasksRepository) GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
			SELECT id, title, description, completed, created_at, completed_at, author_user_id
			FROM todoapp.tasks	
			WHERE id = $1;
		`
	row := r.pool.QueryRow(ctx, query, id)

	var taskModel TaskModel
	err := row.Scan(
		&taskModel.ID,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("task with id='%s': %w", id, core_error.ErrNotFound)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return modelToDomain(taskModel), nil
}
