package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_logger "github.com/MrLaplace19/ToDoList/internal/core/logger"
	core_http_request "github.com/MrLaplace19/ToDoList/internal/core/transport/http/request"
	core_http_response "github.com/MrLaplace19/ToDoList/internal/core/transport/http/response"
)

type GetStatisticsResponse struct{
	CreatedTask int `json:"created_task"`
	CompletedTask int `json:"completed_task"`
	CompletedTaskRate *float64 `json:"completed_task_rate"`
	ArgTime *string `json:"arg_time"`
}

// GetStatistics godoc
// @Summary Получить статистику задач
// @Description Возвращает статистику задач. Фильтры необязательны: user_id — автор, from — начало периода включительно, to — конец периода исключительно; даты в формате YYYY-MM-DD.
// @Tags statistics
// @Produce json
// @Param user_id query int false "Идентификатор автора задач"
// @Param from query string false "Начало периода включительно (YYYY-MM-DD)" format(date)
// @Param to query string false "Конец периода исключительно (YYYY-MM-DD)" format(date)
// @Success 200 {object} GetStatisticsResponse "Успешный ответ"
// @Failure 400 {object} core_http_response.ErrorResponse "Некорректные параметры запроса"
// @Failure 500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /statistics [get]
func (h *StatisticsHTTPHandler) GetStatistics (rw http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, to, from, err := getUserIDFromToQueryParams(r)
	if err!=nil{
		responseHandler.ErrorResponse(
			err,
			"failed to get query params",
		)
		return
	}
	statistics, err := h.statisticsService.GetStatistics(ctx, userID,to,from)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"get statistics service failed",
		)
		return
	}

	response := statisticsResponseFromDomain(statistics)

	responseHandler.JsonResponse(response, http.StatusOK)
}

func statisticsResponseFromDomain(statistics domain.Statistics) GetStatisticsResponse{
	var arg *string
	
	if statistics.ArgTime !=nil{
		duration := statistics.ArgTime.String()
		arg = &duration
	}

	return GetStatisticsResponse{
		CreatedTask: statistics.CreatedTask,
		CompletedTask: statistics.CompletedTask,
		CompletedTaskRate: statistics.CompletedTaskRate,
		ArgTime: arg,	
	}
}

func getUserIDFromToQueryParams(r *http.Request)(*int,*time.Time,*time.Time, error){
	const(
		keyUserID = "user_id"
		keyTo = "to"
		keyFrom = "from"
	)

	id, err := core_http_request.GetIntQueryParams(r, keyUserID)
	if err!=nil{
		return nil,nil,nil, fmt.Errorf("query param `user_id`: %w", err)
	}

	to, err := core_http_request.GetDateQueryParams(r, keyTo)
	if err!=nil{
		return nil,nil,nil, fmt.Errorf("query param `to`: %w", err)
	}

	from, err := core_http_request.GetDateQueryParams(r, keyFrom)
	if err!=nil{
		return nil,nil,nil, fmt.Errorf("query param `from`: %w", err)
	}


	return id,to,from,nil
}