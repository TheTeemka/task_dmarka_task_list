package models

type TaskFilter struct {
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
}

func (t *TaskFilter) Validate() error {
	return validate.Struct(t)
}
