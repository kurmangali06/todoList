package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"todolist/internal/core/domain"
	core_error "todolist/internal/core/errors"
	core_postgres_pool "todolist/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO todoapp.tasks  (title, description, completed, created_at, completed_at, author_user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;
		`
	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.AuthorUserID,
	)
	var taskModel TaskModel

	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserID,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf(
				"%v: user with id='%s': %w",
				err,
				task.AuthorUserID,
				core_error.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := modelToDomain(taskModel)
	return taskDomain, nil
}
