package domain

import (
	"time"
)

type Statistics struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64       // процент выполненных задач (nil если задач нет)
	TasksAverageCompletionTime *time.Duration // среднее время выполнения (nil если нет выполненных)
}

func CreateStatistics(tasks []Task) Statistics {
	if len(tasks) == 0 {
		return Statistics{
			TasksCreated:               0,
			TasksCompleted:             0,
			TasksCompletedRate:         nil,
			TasksAverageCompletionTime: nil,
		}
	}
	tasksCreated := len(tasks)

	tasksCompleted := 0
	var totalCompletionDuration time.Duration
	for _, task := range tasks {
		if task.CompletedAt != nil {
			tasksCompleted++
		}
		completionDuration := task.CompletionDuration()

		if completionDuration != nil {
			totalCompletionDuration += *completionDuration
		}
	}
	tasksCompletedRate := float64(tasksCompleted) / float64(tasksCreated) * 100

	var tasksAverageCompletionTime *time.Duration
	if tasksCompleted > 0 && totalCompletionDuration != 0 {
		avg := totalCompletionDuration / time.Duration(tasksCompleted)

		tasksAverageCompletionTime = &avg
	}

	return NewStatistics(
		tasksCreated,
		tasksCompleted,
		&tasksCompletedRate,
		tasksAverageCompletionTime,
	)
}

func NewStatistics(
	tasksCreated int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
