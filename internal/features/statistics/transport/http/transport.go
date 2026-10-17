package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_http_server "github.com/MrLaplace19/ToDoList/internal/core/transport/http/server"
)


type StatisticsHTTPHandler struct{
	statisticsService StatisticsService
}

type StatisticsService interface{
	GetStatistics(
		ctx context.Context,
		id *int,
		to *time.Time,
		from *time.Time,
	)(domain.Statistics, error)
}

func NewStatisticsHTTPHandler (
	statisticsService StatisticsService,
) *StatisticsHTTPHandler{
	return &StatisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Router{
	return []core_http_server.Router{
		{
			Method: http.MethodGet,
			Path: "/statistics",
			Handler: h.GetStatistics,
		},
	}
}