package tasks_service

import (
	"context"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
)

func (s *TasksService) GetTasks(
		ctx context.Context,
		limit *int,
		offset *int,
	)([]domain.Task, error){
		if limit !=nil && *limit <0 {
			return []domain.Task{}, fmt.Errorf(
				"'limit'=%d must be >0: %w",
				*limit,
				core_errors.ErrInvalidArgument,
			)
		}

		if offset != nil && *offset < 0{
			return []domain.Task{}, fmt.Errorf(
				"'offset'=%d must be >0: %w",
				*offset,
				core_errors.ErrInvalidArgument,
			)
		}

		tasksDomain, err := s.taskRepository.GetTasks(ctx, limit, offset)
		if err!=nil{
			return []domain.Task{},fmt.Errorf("service get users failde: %w", err)
		}

		return tasksDomain, nil
	}