package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"todolist/internal/core/domain"
	core_error "todolist/internal/core/errors"
	core_postgres_pool "todolist/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) PatchTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE todoapp.tasks
		SET 
		    title = $1, 
		    description = $2, 
		    completed = $3, 
		    completed_at = $4,
		    version = version + 1
		WHERE id = $5 AND version = $6
		
		RETURNING 
			id, 
			version, 
			title, 
			description, 
			completed, 
			created_at, 
			completed_at, 
			author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.ID,
		task.Version,
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
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%s' concurrently accessed: %w",
				task.ID,
				core_error.ErrConflict,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}
	taskDomain := modelToDomain(taskModel)

	return taskDomain, nil
}
