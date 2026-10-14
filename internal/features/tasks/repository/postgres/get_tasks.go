package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

func (r *TasksRepository) GetTasks(
		ctx context.Context,
		limit *int,
		offset *int,
	)([]domain.Task, error){
		ctx, cancel := context.WithTimeout(ctx, r.pool.OptionalTimeOut())
		defer cancel()

		query := `
		SELECT id,version,title,description,created_at,completed,completed_at,author_id
		FROM todoapp.tasks
		ORDER BY id ASC 
		LIMIT $1
		OFFSET $2;
		`

		rows, err := r.pool.Query(ctx, query, limit, offset)
		if err!=nil {
			return nil, fmt.Errorf("select tasks: %w", err)
		}

		defer rows.Close()

		var taskModels []TaskModel
		for rows.Next(){
			var taskModel TaskModel
			err := rows.Scan(
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
				return nil, fmt.Errorf("scan tasks: %w", err)
			}
			taskModels = append(taskModels, taskModel)
		}

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("next rows: %w", err)
		}

		tasksDomain := tasksDomainFromModels(taskModels...)
		return tasksDomain, nil
	}