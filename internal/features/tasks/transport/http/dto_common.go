package tasks_transport_http

import (
	"time"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

type TaskDTOResponse struct {
	ID           uuid.UUID  `json:"id"             example:"15"`
	Version      int        `json:"version"        example:"3"`
	Title        string     `json:"title"          example:"Домашка"`
	Description  *string    `json:"description"    example:"Сделать до четверга домашнее задание по математике"`
	Completed    bool       `json:"completed"      example:"false"`
	CreatedAt    time.Time  `json:"created_at"     example:"2026-02-26T10:30:00Z"`
	CompletedAt  *time.Time `json:"completed_at"   example:"null"`
	AuthorUserID uuid.UUID  `json:"author_user_id" example:"5"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func taskDTOsFromDomains(tasks []domain.Task) []TaskDTOResponse {
	dtos := make([]TaskDTOResponse, len(tasks))

	for i, task := range tasks {
		dtos[i] = taskDTOFromDomain(task)
	}

	return dtos
}
