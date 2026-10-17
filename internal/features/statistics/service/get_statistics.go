package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/MrLaplace19/ToDoList/internal/core/domain"
	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
		ctx context.Context,
		id *int,
		to *time.Time,
		from *time.Time,
	)(domain.Statistics, error){
		if from !=nil && to != nil{
			if to.Before(*from) || to.Equal(*from){
				return domain.Statistics{}, fmt.Errorf(
					"`to` must be after `from`: %w",
					core_errors.ErrInvalidArgument,
				)
			}
		}

		tasks, err := s.statisticsRepository.GetTasks(ctx, id, to, from)
		if err!=nil{
			return domain.Statistics{}, fmt.Errorf(
				"get tasks from repository: %w",
				err,
			)
		}

		statistics := calcStatistics(tasks)

		return statistics, nil
}



func calcStatistics(domains []domain.Task) domain.Statistics{
	if len(domains) == 0{
		return domain.Statistics{
			CreatedTask: 0,
			CompletedTask: 0,
			CompletedTaskRate: nil,
			ArgTime: nil,
		}
	}
	createdTask := len(domains)
	completedTask := 0
	var totalCompletionDuration time.Duration
	for _,v := range domains{
		if v.Completed{
			completedTask++
		}

		completionDuration := v.CompletedDuration()
		if completionDuration!=nil{
			totalCompletionDuration+=*completionDuration
		}
	}

	taskCompletedRate := float64(completedTask) / float64(createdTask) * 100

	var average *time.Duration

	if completedTask > 0 && totalCompletionDuration !=0{
		avg := totalCompletionDuration /time.Duration(completedTask)
		average = &avg
	}

	return domain.Statistics{
		CreatedTask: createdTask,
		CompletedTask: completedTask,
		CompletedTaskRate: &taskCompletedRate,
		ArgTime: average,
	}
}	