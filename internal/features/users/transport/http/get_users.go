package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
	core_http_utils "github.com/MrLaplace19/ToDoList/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

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

	limit, err := core_http_utils.GetIntQueryParams(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query params %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParams(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query params %w", err)
	}

	return limit, offset, nil
}
