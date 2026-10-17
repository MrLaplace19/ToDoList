package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

func (r *StatisticsRepository) GetTasks(
		ctx context.Context,
		id *int,
		to *time.Time,
		from *time.Time,
	)([]domain.Task, error){
		ctx, cancel := context.WithTimeout(ctx, r.pool.OptionalTimeOut())
		defer cancel()

		var queryBuilder strings.Builder
		queryBuilder.WriteString(`
		SELECT id,version,title,description,completed,completed_at,created_at,author_id
		FROM todoapp.tasks
		`)

		args := []any{}
		conditions :=[]string{}
		if id!=nil{
			conditions = append(conditions, fmt.Sprintf("author_id=$%d", len(args)+1))
			args = append(args, id)
		}
		if from!=nil{
			conditions = append(conditions, fmt.Sprintf("created_at >=$%d", len(args)+1))
			args = append(args, from)
		}
		if to!=nil{
			conditions = append(conditions, fmt.Sprintf("created_at <$%d", len(args)+1))
			args = append(args, from)
		}

		if len(conditions) != 0{
			queryBuilder.WriteString((" WHERE " + strings.Join(conditions, " AND ")))
		}

		queryBuilder.WriteString(" ORDER BY id ASC")

		rows, err := r.pool.Query(ctx, queryBuilder.String(), args...)
		if err != nil{
			return nil, fmt.Errorf(
				"select tasks: %w",
				err,
			)
		}
		defer rows.Close()

		models := []TaskModel{}
		for rows.Next(){
			var model TaskModel
			err := rows.Scan(
				&model.ID,
				&model.Version,
				&model.Title,
				&model.Description,
				&model.Completed,
				&model.CompletedAt,
				&model.CreatedAt,
				&model.AuthorUserID,
			)

			if err != nil{
				return nil, fmt.Errorf(
					"scan rows: %w",
					err,
				)
			}

			models = append(models, model)
		}

		domains := tasksDomainFromModels(models...)
		return domains, nil
	}