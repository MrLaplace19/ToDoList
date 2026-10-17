package domain

import "time"

type Statistics struct{
	CreatedTask int
	CompletedTask int
	CompletedTaskRate *float64
	ArgTime *time.Duration
}