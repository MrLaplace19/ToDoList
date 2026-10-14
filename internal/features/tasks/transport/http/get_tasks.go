package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_request "github.com/MrLaplace19/ToDoList/internal/core/transport/http/request"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
)

type GetTasksResponse []DTOTaskResponse

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit,offset,err := getLimitOffsetQueryParams(r)
	if err != nil{
		responseHandler.ErrorResponse(
			err,
			"failed to get 'limit/offset': %w",
		)
		return
	}

	tasks, err := h.tasksService.GetTasks(ctx, limit, offset)
	if err!=nil{
		responseHandler.ErrorResponse(
			err,
			"failed to get tasks: %w",
		)
		return
	}

	response := GetTasksResponse(tasksDTOFromDomains(tasks...))

	responseHandler.JsonResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int,*int, error){
	const (
		keyLimit = "limit"
		keyOffset = "offset"
	)
	limit, err := core_http_request.GetIntQueryParams(r, keyLimit)
	if err !=nil {
		return nil,nil,fmt.Errorf("error get limit: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParams(r, keyOffset)
	if err!=nil{
		return nil,nil,fmt.Errorf("error get limit: %w", err)
	}

	return limit, offset, nil
}