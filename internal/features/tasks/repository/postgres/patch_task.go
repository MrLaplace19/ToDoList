package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
	core_postgres_pool "github.com/MrLaplace19/ToDoList/internal/core/repository/postgres/pool"
)

func(r *TasksRepository) PatchTask(
		ctx context.Context,
		task domain.Task,
		id int,
	)(domain.Task, error){
		ctx, cancel := context.WithTimeout(ctx, r.pool.OptionalTimeOut())
		defer cancel()

		query := `
		UPDATE todoapp.tasks
		SET
			version=version+1,
			title=$1,
			description=$2,
			completed=$3,
			completed_at=$4
		WHERE id=$5 AND version=$6
		RETURNING 
			id,
			version,
			title,
			description,
			created_at,
			completed,
			completed_at,
			author_id;
		`

		row := r.pool.QueryRow(
			ctx,
			query,
			task.Title,
			task.Description,
			task.Completed,
			task.CompletedAt,
			id,
			task.Version,
		)

		var taskModel TaskModel

		err := row.Scan(
			&taskModel.ID,
			&taskModel.Version,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.CreatedAt,
			&taskModel.Completed,
			&taskModel.CompletedAt,
			&taskModel.AuthorUserID,
		)

		if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"user with id='%v' concurently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

			return domain.Task{}, fmt.Errorf("user scan to patch: %w", err)
		}

		taskDomain := domain.Task{
			ID: taskModel.ID,
			Version: taskModel.Version,
			Title: taskModel.Title,
			Description: taskModel.Description,
			Completed: taskModel.Completed,
			CompletedAt: taskModel.CompletedAt,
			CreatedAt: taskModel.CreatedAt,
			AuthorUserID: taskModel.AuthorUserID,
		}

		return taskDomain, nil
	}