package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_request "github.com/MrLaplace19/ToDoList/internal/core/transport/http/request"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
)

// DeleteTask godoc
// @Summary Удалить задачу
// @Description Удаляет задачу по идентификатору.
// @Tags tasks
// @Produce json
// @Param id path int true "Идентификатор"
// @Success 204 "Удалено, тело ответа отсутствует"
// @Failure 400 {object} core_http_response.ErrorResponse "Некорректные параметры запроса"
// @Failure 404 {object} core_http_response.ErrorResponse "Объект не найден"
// @Failure 500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `id`",
		)
		return
	}

	if err := h.tasksService.DeleteTask(ctx, id); err != nil{
		responseHandler.ErrorResponse(
			err,
			"failed to delete task",
		)
		return
	}

	responseHandler.NoContentResponse()
}