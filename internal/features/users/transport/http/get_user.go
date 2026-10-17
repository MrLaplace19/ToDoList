package users_transport_http

import (
	"net/http"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_request "github.com/MrLaplace19/ToDoList/internal/core/transport/http/request"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

// GetUser godoc
// @Summary Получить пользователя
// @Description Возвращает пользователя по идентификатору.
// @Tags users
// @Produce json
// @Param id path int true "Идентификатор"
// @Success 200 {object} GetUserResponse "Успешный ответ"
// @Failure 400 {object} core_http_response.ErrorResponse "Некорректные параметры запроса"
// @Failure 404 {object} core_http_response.ErrorResponse "Объект не найден"
// @Failure 500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /users/{id} [get]
func (h *UserHttpHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)
		return
	}

	user, err := h.userService.GetUser(ctx, userID)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JsonResponse(response, http.StatusOK)
}
