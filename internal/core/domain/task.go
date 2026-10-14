package domain

import (
	"fmt"
	"time"

	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
)

type Task struct{
	ID int
	Version int

	Title string
	Description *string
	Completed bool
	CompletedAt *time.Time
	CreatedAt time.Time
	AuthorUserID int
}

func NewTask(
	id, version int,
	title string,
	description *string,
	completed bool,
	completedAt *time.Time,
	createdAt time.Time,
	authorID int,
) Task{
	return Task{
		ID: id,
		Version: version,
		Title: title,
		Description: description,
		Completed: completed,
		CompletedAt: completedAt,
		CreatedAt: createdAt,
		AuthorUserID: authorID,
	}
}

func NewTaskUninitialized(title string, description *string, authorID int) Task{
	return NewTask(
		UninitializedID,
		UninitializedVersion,
		title,
		description,
		false,
		nil,
		time.Now(),
		authorID,
	)
}

func (t *Task) Validate() error{

	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen>100{
		return fmt.Errorf("invalid 'title' len %d: %w", titleLen, core_errors.ErrInvalidArgument)
	}

	if t.Description != nil{
		descriptionLen := len([]rune(*t.Description))
		if descriptionLen < 1 || descriptionLen > 1000{
			return fmt.Errorf("invalid 'description' len %d: %w", descriptionLen, core_errors.ErrInvalidArgument)
		}
	}

	if t.Completed {
		if t.CompletedAt == nil{
			return fmt.Errorf("CompletedAT can`t be 'nil' if 'Completed == true': %w", core_errors.ErrInvalidArgument)
		}

		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("CompletedAT can`t be > 'CreatedAT': %w", core_errors.ErrInvalidArgument)
		}
	}else{
		if t.CompletedAt != nil {
			return fmt.Errorf("CompletedAt must be 'nil' if 'Completed == false': %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

type PatchTask struct{
	Title Nullable[string]
	Description Nullable[string]
	Completed Nullable[bool]
}

func NewPatchTask(
	title Nullable[string],
	description Nullable[string],
	completed Nullable[bool],
) PatchTask{
	return PatchTask{
		Title: title,
		Description: description,
		Completed: completed,
	}
}

func (p *PatchTask) Validate() error{
	if p.Title.Set && p.Title.Value == nil{
		return fmt.Errorf(
			"`Title` can`t be null: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Completed.Set && p.Completed.Value == nil{
		return fmt.Errorf(
			"`Completed` can`t be null: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (t *Task) ApplyPatch(patch PatchTask) error{
	if err := patch.Validate(); err !=nil{
		return fmt.Errorf(
			"validate patch failed: %w",
			err,
		)
	}

	tmp := *t

	if patch.Title.Set{
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set{
		tmp.Description = patch.Description.Value
	}

	if patch.Completed.Set{
		tmp.Completed = *patch.Completed.Value

		if tmp.Completed{
			completedAt := time.Now() 
			tmp.CompletedAt = &completedAt
		}else{
			tmp.CompletedAt = nil
		}
	}

	if err := tmp.Validate(); err !=nil{
		return fmt.Errorf(
			"validate patch task failed: %w",
			err,
		)
	}

	*t = tmp

	return nil
}
