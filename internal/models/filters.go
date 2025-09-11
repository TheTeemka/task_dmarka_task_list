package models

type TaskFilter struct {
	Limit     *int    `validate:"omitempty,gte=1,lte=100"`
	Offset    *int    `validate:"omitempty,gte=0"`
	TagIDs    []int64 `validate:"omitempty,dive,gte=1"`
	Status    *int    `validate:"omitempty,gte=0,lte=3"`
	Priority  *int    `validate:"omitempty,gte=0,lte=4"`
	SortBy    *string `validate:"omitempty,oneof=id title status priority start_date due_date created_date"`
	SortOrder *string `validate:"omitempty,oneof=asc desc"`
}

func (t *TaskFilter) Validate() error {
	return validate.Struct(t)
}
