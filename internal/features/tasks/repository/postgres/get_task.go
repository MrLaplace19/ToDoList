package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
	core_postgres_pool "github.com/MrLaplace19/ToDoList/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) GetTask(
		ctx context.Context,
		id int,
	)(domain.Task, error){
		ctx, cancel := context.WithTimeout(ctx, r.pool.OptionalTimeOut())
		defer cancel()

		query := `
		SELECT id,version,title,description,created_at,completed,completed_at,author_id
		FROM todoapp.tasks
		WHERE id=$1;
		`

		row := r.pool.QueryRow(ctx, query, id)
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
				"user with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

		task := domain.Task{
			ID: taskModel.ID,
			Version: taskModel.Version,
			Title: taskModel.Title,
			Description: taskModel.Description,
			Completed: taskModel.Completed,
			CompletedAt: taskModel.CompletedAt,
			CreatedAt: taskModel.CreatedAt,
			AuthorUserID: taskModel.AuthorUserID,
		}

		return task, nil
	}