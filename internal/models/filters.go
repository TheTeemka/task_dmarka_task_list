package models

import "time"

type TaskFilterModel struct {
	Limit  *int
	Offset *int

	TagIDs  []int64
	TagsNot bool

	Status    []int
	StatusNot bool

	Priority    []int
	PriorityNot bool

	Search *string

	SortBy    *string
	SortOrder *string

	Before *time.Time
}

func (t *TaskFilterModel) ToDTO() TaskFilterDTO {
	dto := TaskFilterDTO{
		Limit:       t.Limit,
		Offset:      t.Offset,
		TagIDs:      t.TagIDs,
		TagsNot:     t.TagsNot,
		Status:      t.Status,
		StatusNot:   t.StatusNot,
		Priority:    t.Priority,
		PriorityNot: t.PriorityNot,
		Search:      t.Search,
		SortBy:      t.SortBy,
		SortOrder:   t.SortOrder,
		Before:      nil,
	}
	if t.Before != nil {
		beforeStr := t.Before.Format(time.RFC3339)
		dto.Before = &beforeStr
	}
	return dto
}

type TaskFilterDTO struct {
	Limit     *int    `validate:"omitempty,gte=1,lte=100"`
	Offset    *int    `validate:"omitempty,gte=0"`
	TagIDs    []int64 `validate:"omitempty,dive,gte=1"`
	TagsNot   bool    `validate:"omitempty"`
	Status    []int   `validate:"omitempty,dive,gte=1,lte=4"`
	StatusNot bool    `validate:"omitempty"`

	Priority    []int `validate:"omitempty,dive,gte=1,lte=4"`
	PriorityNot bool  `validate:"omitempty"`

	Search *string `validate:"omitempty"`

	SortBy    *string `validate:"omitempty,oneof=id title status priority due_date completed_date created_at duedate "`
	SortOrder *string `validate:"omitempty,oneof=asc desc"`

	Before *string `validate:"omitempty"`
}

func (t *TaskFilterDTO) ToModel() *TaskFilterModel {
	model := &TaskFilterModel{
		Limit:       t.Limit,
		Offset:      t.Offset,
		TagIDs:      t.TagIDs,
		TagsNot:     t.TagsNot,
		Status:      t.Status,
		StatusNot:   t.StatusNot,
		Priority:    t.Priority,
		PriorityNot: t.PriorityNot,
		Search:      t.Search,
		SortBy:      t.SortBy,
		SortOrder:   t.SortOrder,
		Before:      nil,
	}
	if t.Before != nil {
		beforeTime, err := time.Parse(time.RFC3339, *t.Before)
		if err == nil {
			model.Before = &beforeTime
		}
	}
	return model
}

func (t *TaskFilterDTO) Validate() error {
	return validate.Struct(t)
}
