package tasks_service

import (
	"context"

	"github.com/Mommsent/todoapp-Studying.git/internal/core/domain"
)

type TaskService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)

	GetTasks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Task, error)

	GetTask(
		ctx context.Context,
		id int,
	) (domain.Task, error)

	DeleteTask(
		ctx context.Context,
		id int,
	) error

	PatchTask(
		ctx context.Context,
		id int,
		task domain.Task,
	) (domain.Task, error)
}

func NewTasksService(
	tasksRepo TasksRepository,
) *TaskService {
	return &TaskService{
		tasksRepository: tasksRepo,
	}
}
