package tasks_postgres_repository

import (
	"context"
	"fmt"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

func (r *TasksRepository) GetTasks(
	ctx context.Context,
	userID *uuid.UUID,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
		FROM todoapp.tasks
		%s
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2;
		`
	args := []any{limit, offset}
	if userID != nil {
		query = fmt.Sprintf(query, "WHERE author_user_id = $3")
		args = append(args, userID)
	} else {
		query = fmt.Sprintf(query, "")
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)
	}
	defer rows.Close()
	var tasksModel []TaskModel

	for rows.Next() {
		var taskModel TaskModel
		err := rows.Scan(
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
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasksModel = append(tasksModel, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows tasks: %w", err)
	}
	taskDomains := modelsToDomains(tasksModel)
	return taskDomains, nil
}
