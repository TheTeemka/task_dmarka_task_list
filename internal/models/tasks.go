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
	dto := &TaskDTO{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Priority:    t.Priority,
	}

	// Convert time.Time to RFC3339 strings, empty string for zero time
	if !t.DueDate.IsZero() {
		dto.DueDate = t.DueDate.Format(time.RFC3339)
	}
	if !t.CompletedDate.IsZero() {
		dto.CompletedDate = t.CompletedDate.Format(time.RFC3339)
	}
	if !t.CreatedDate.IsZero() {
		dto.CreatedDate = t.CreatedDate.Format(time.RFC3339)
	}

	return dto
}

type TaskDTO struct {
	ID            int64  `json:"id" validate:"omitempty,min=1"`
	Title         string `json:"title" validate:"required,min=1,max=100"`
	Description   string `json:"description" validate:"omitempty,max=500"`
	Status        int    `json:"status" validate:"min=0,max=4"`
	Priority      int    `json:"priority" validate:"min=0,max=4"`
	DueDate       string `json:"due_date" validate:"omitempty"`
	CompletedDate string `json:"completed_date" validate:"omitempty"`
	CreatedDate   string `json:"created_at" validate:"omitempty"`
}

func (t *TaskDTO) Validate() error {
	return validate.Struct(t)
}

func (t *TaskDTO) ToModel() *TaskModel {
	model := &TaskModel{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Priority:    t.Priority,
	}

	// Parse RFC3339 strings to time.Time, zero time for empty strings
	if t.DueDate != "" {
		if parsed, err := time.Parse(time.RFC3339, t.DueDate); err == nil {
			model.DueDate = parsed
		}
	}
	if t.CompletedDate != "" {
		if parsed, err := time.Parse(time.RFC3339, t.CompletedDate); err == nil {
			model.CompletedDate = parsed
		}
	}
	if t.CreatedDate != "" {
		if parsed, err := time.Parse(time.RFC3339, t.CreatedDate); err == nil {
			model.CreatedDate = parsed
		}
	}

	return model
}

// Status represents the task status as a string enum
type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusCancelled  Status = "cancelled"
)
