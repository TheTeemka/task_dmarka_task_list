package models

import "time"

type TaskModel struct {
	ID          int64
	Title       string
	Description string
	Status      int64
	PriorityID  int64

	StartDate time.Time
	DueDate   time.Time
}

func (t *TaskModel) ToDTO() *TaskDTO {
	return &TaskDTO{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		PriorityID:  t.PriorityID,
		StartDate:   t.StartDate,
		DueDate:     t.DueDate,
	}
}

type TaskDTO struct {
	ID          int64
	Title       string
	Description string
	Status      int64
	PriorityID  int64

	StartDate time.Time
	DueDate   time.Time

	CreatedDate time.Time
}

func (t *TaskDTO) ToModel() *TaskModel {
	return &TaskModel{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		PriorityID:  t.PriorityID,
		StartDate:   t.StartDate,
		DueDate:     t.DueDate,
	}
}
