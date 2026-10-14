package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_request "github.com/MrLaplace19/ToDoList/internal/core/transport/http/request"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
	core_http_types "github.com/MrLaplace19/ToDoList/internal/core/transport/http/types"
)

type PatchTaskRequest struct{
	Title core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed core_http_types.Nullable[bool]	`json:"completed"`
}

type PatchTaskResponse DTOTaskResponse

func (p *PatchTaskRequest) Validate() error{
	if p.Title.Set{
		if p.Title.Value == nil {
			return fmt.Errorf("`Title` can`t be null")
		}

		titleLen := len([]rune(*p.Title.Value))
		if titleLen < 1 || titleLen>100{
			return fmt.Errorf("`Title` must be between 1 and 100")
		}
	}

	if p.Description.Set{
		if p.Description.Value !=nil {
			descriptionLen := len([]rune(*p.Description.Value))

			if descriptionLen <1 || descriptionLen >1000{
				return fmt.Errorf("`Description` must be between 1 and 1000")
			}
		}
	}

	if p.Completed.Set{
		if p.Completed.Value == nil{
			return fmt.Errorf("`Completed` can`t be null")
		}
	}

	return nil
}

func (h *TasksHTTPHandler) PatchTask (rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err!=nil{
		responseHandler.ErrorResponse(
			err,
			"failed to get 'id'",
		)
		return
	}

	var request PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err !=nil{
		responseHandler.ErrorResponse(
			err,
			"validate task failed: %w",
		)
		return
	}

	taskPatch := patchTaskFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskPatch, id)
	if err!=nil {
		responseHandler.ErrorResponse(
			err,
			"failed task patch: %w",
		)
		return
	}

	response := PatchTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JsonResponse(response, http.StatusOK)
}


func patchTaskFromRequest(request PatchTaskRequest) domain.PatchTask {
	return domain.NewPatchTask(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}