package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_request "github.com/MrLaplace19/ToDoList/internal/core/transport/http/request"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
)

func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'id'",
		)
		return
	}

	if err := h.tasksService.DeleteTask(ctx, id); err != nil{
		responseHandler.ErrorResponse(
			err,
			"failed to delete task: %w",
		)
		return
	}

	responseHandler.NoContentResponse()
}