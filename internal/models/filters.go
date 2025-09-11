package models

type TaskFilter struct {
	Limit  *int
	Offset *int

	TagIDs     []int64
	PriorityID *int64

	SortBy    *string
	SortOrder *int
}
