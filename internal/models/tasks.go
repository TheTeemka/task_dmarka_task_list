package models

import (
	"time"
)

type TaskModel struct {
	ID            int64
	Title         string
	Description   string
	Status        int
	Priority      int
	DueDate       time.Time
	CompletedDate time.Time
	CreatedDate   time.Time
}

func (t *TaskModel) ToDTO() *TaskDTO {
	return &TaskDTO{
		ID:            t.ID,
		Title:         t.Title,
		Description:   t.Description,
		Status:        t.Status,
		Priority:      t.Priority,
		DueDate:       t.DueDate,
		CompletedDate: t.CompletedDate,
		CreatedDate:   t.CreatedDate,
	}
}

type TaskDTO struct {
	ID            int64     `validate:"omitempty,min=1"`
	Title         string    `validate:"required,min=1,max=100"`
	Description   string    `validate:"omitempty,max=500"`
	Status        int       `validate:"min=0,max=4"`
	Priority      int       `validate:"min=0,max=4"`
	DueDate       time.Time `validate:"omitempty"`
	CompletedDate time.Time `validate:"omitempty"`
	CreatedDate   time.Time `validate:"omitempty"`
}

func (t *TaskDTO) Validate() error {
	return validate.Struct(t)
}

func (t *TaskDTO) ToModel() *TaskModel {
	return &TaskModel{
		ID:            t.ID,
		Title:         t.Title,
		Description:   t.Description,
		Status:        t.Status,
		Priority:      t.Priority,
		DueDate:       t.DueDate,
		CompletedDate: t.CompletedDate,
		CreatedDate:   t.CreatedDate,
	}
}

// Status represents the task status as a string enum
type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusCancelled  Status = "cancelled"
)
