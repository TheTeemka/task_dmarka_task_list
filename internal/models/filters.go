package models

type TaskFilter struct {
	Limit     *int    `validate:"omitempty,min=1,max=100"`
	Offset    *int    `validate:"omitempty,min=0"`
	TagIDs    []int64 `validate:"omitempty,dive,min=1"`
	Priority  *int    `validate:"omitempty,oneof=low medium high urgent"`
	SortBy    *string `validate:"omitempty,oneof=title status priority created_date"`
	SortOrder *int    `validate:"omitempty,oneof=0 1"` // 0 for ASC, 1 for DESC
}

func (t *TaskFilter) Validate() error {
	return validate.Struct(t)
}
