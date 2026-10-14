package tasks_postgres_repository

import (
	"time"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

type TaskModel struct{
	ID int
	Version int

	Title string
	Description *string
	Completed bool
	CompletedAt *time.Time
	CreatedAt time.Time
	AuthorUserID int
}

func tasksDomainFromModels(models ...TaskModel) []domain.Task{
	domains := make([]domain.Task, len(models))

	for i,v := range models{
		domains[i] = domain.Task{
			ID: v.ID,
			Version: v.Version,
			Title: v.Title,
			Description: v.Description,
			Completed: v.Completed,
			CompletedAt: v.CompletedAt,
			CreatedAt: v.CreatedAt,
			AuthorUserID: v.AuthorUserID,
		}
	}

	return domains
}