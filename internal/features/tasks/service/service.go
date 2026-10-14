package tasks_service

import (
	"context"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)


type TasksService struct{
	taskRepository TasksRepository
}

type TasksRepository interface{
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
	GetTasks(
		ctx context.Context,
		limit *int,
		offset *int,
	)([]domain.Task, error)
	GetTask(
		ctx context.Context,
		id int,
	)(domain.Task, error)
	DeleteTask(
		ctx context.Context,
		id int,
	)error
	PatchTask(
		ctx context.Context,
		task domain.Task,
		id int,
	)(domain.Task, error)
}

func NewTasksService(
	taskRepository TasksRepository,
) *TasksService{
	return &TasksService{
		taskRepository: taskRepository,
	}
}