package tasks_transport_http

import (
	"time"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)


type DTOTaskResponse struct{
	ID int `json:"id"`
	Version int `json:"version"`

	Title string `json:"title"`
	Description *string `json:"description"`
	Completed bool	`json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`

	AuthorUserID int `json:"author_id"`
}

func taskDTOFromDomain(domain domain.Task) DTOTaskResponse{
	return DTOTaskResponse{
		ID: domain.ID,
		Version: domain.Version,

		Title: domain.Title,
		Description: domain.Description,
		Completed: domain.Completed,
		CreatedAt: domain.CreatedAt,
		CompletedAt: domain.CompletedAt,

		AuthorUserID: domain.AuthorUserID,
	}
}

func tasksDTOFromDomains(domains ...domain.Task) []DTOTaskResponse{
	dtoTasksResponse := make([]DTOTaskResponse, len(domains))
	for i,v := range domains{
		dtoTasksResponse[i] = taskDTOFromDomain(v)
	}

	return dtoTasksResponse
}