package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_request "github.com/MrLaplace19/ToDoList/internal/core/transport/http/request"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
)

type GetUsersResponse []UserDTOResponse

// GetUsers godoc
// @Summary Получить пользователей
// @Description Возвращает список пользователей с необязательной пагинацией.
// @Tags users
// @Produce json
// @Param limit query int false "Максимальное количество записей" minimum(0)
// @Param offset query int false "Количество пропускаемых записей" minimum(0)
// @Success 200 {object} GetUsersResponse "Успешный ответ"
// @Failure 400 {object} core_http_response.ErrorResponse "Некорректные параметры запроса"
// @Failure 500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /users [get]
func (h *UserHttpHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset")
		return
	}
	userDomains, err := h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
	}

	response := GetUsersResponse(usersDTOFromDomains(userDomains))
	responseHandler.JsonResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {

	const (
		keyLimit = "limit"
		keyOffset = "offset"
	)

	limit, err := core_http_request.GetIntQueryParams(r, keyLimit)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query params %w", err)
	}

	offset, err := core_http_request.GetIntQueryParams(r, keyOffset)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query params %w", err)
	}

	return limit, offset, nil
}
