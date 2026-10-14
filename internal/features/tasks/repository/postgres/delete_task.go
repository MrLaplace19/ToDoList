package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(
		ctx context.Context,
		id int,
	)error{
		ctx, cancel := context.WithTimeout(ctx, r.pool.OptionalTimeOut())
		defer cancel()

		query := `
		DELETE FROM todoapp.tasks
		WHERE id=$1;
		`

		comTag, err := r.pool.Exec(ctx, query, id)
		if err!=nil{
			return fmt.Errorf("exec query: %w", err)
		}
		if comTag.RowsAffected() == 0 {
			return fmt.Errorf("user with id='%d': %w", id, core_errors.ErrNotFound)
		}

		return nil
	}