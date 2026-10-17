package statistics_service

import (
	"context"
	"time"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
)

type StatisticsService struct{
	statisticsRepository StatisticsRepository
}

type StatisticsRepository interface{
	GetTasks(
		ctx context.Context,
		id *int,
		to *time.Time,
		from *time.Time,
	)([]domain.Task, error)
}

func NewStatisticsService(
	statisticsRepository StatisticsRepository,
) *StatisticsService{
	return &StatisticsService{
		statisticsRepository: statisticsRepository,
	}
}