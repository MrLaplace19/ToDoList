package tasks_service

import (
	"context"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

func (s *TasksService) GetTask(
		ctx context.Context,
		id int,
	)(domain.Task, error){
		task, err := s.taskRepository.GetTask(ctx, id)
		if err !=nil{
			return domain.Task{}, fmt.Errorf(
				"service get task: %w",
				err,
			)
		}

		return task, nil
	}