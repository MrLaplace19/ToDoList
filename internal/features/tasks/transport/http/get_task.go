package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_request "github.com/MrLaplace19/ToDoList/internal/core/transport/http/request"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
)

type GetTaskResponse DTOTaskResponse

// GetTask godoc
// @Summary Получить задачу
// @Description Возвращает задачу по идентификатору.
// @Tags tasks
// @Produce json
// @Param id path int true "Идентификатор"
// @Success 200 {object} GetTaskResponse "Успешный ответ"
// @Failure 400 {object} core_http_response.ErrorResponse "Некорректные параметры запроса"
// @Failure 404 {object} core_http_response.ErrorResponse "Объект не найден"
// @Failure 500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err !=nil{
		responseHandler.ErrorResponse(
			err,
			"failed to get `id`",
		)
		return
	}

	task,err := h.tasksService.GetTask(ctx, id)
	if err!=nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task",
		)
		return
	} 

	response := GetTaskResponse(taskDTOFromDomain(task))
	responseHandler.JsonResponse(response, http.StatusOK)
}