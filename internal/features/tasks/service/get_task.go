package tasks_service

import (
	"context"
	"fmt"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
)

func (s *TaskService) GetTask(ctx context.Context, id int) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}

	return task, nil
}
