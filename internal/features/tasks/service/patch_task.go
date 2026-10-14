package tasks_service

import (
	"context"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

func (s *TasksService) PatchTask(
		ctx context.Context,
		patch domain.PatchTask,
		id int,
	)(domain.Task, error){
		task, err := s.taskRepository.GetTask(ctx, id)
		if err!=nil{
			return domain.Task{}, fmt.Errorf(
				"Task get task: %w",
				err,
			)
		}

		if err := task.ApplyPatch(patch); err!=nil{
			return domain.Task{}, fmt.Errorf(
				"apply task patch: %w", 
				err,
			)
		}

		taskDomain, err := s.taskRepository.PatchTask(ctx, task, id)
		if err!=nil{
			return domain.Task{}, fmt.Errorf(
				"patch task: %w",
				err,
			)
		}

		return taskDomain, nil
	}