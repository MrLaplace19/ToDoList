package tasks_service

import (
	"context"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

func (s *TasksService) CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error){
		if err := task.Validate(); err != nil {
			return domain.Task{}, fmt.Errorf("Task not be validate: %w", err)
		}

		newTask, err := s.taskRepository.CreateTask(ctx, task)

		if err !=nil{
			return domain.Task{}, fmt.Errorf("create task: %w", err)
		}

		return newTask, nil
	}