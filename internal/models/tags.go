package models

type TagModel struct {
	ID    int64
	Name  string
	Color string
}

func (t *TagModel) ToDTO() *TagDTO {
	return &TagDTO{
		ID:    t.ID,
		Name:  t.Name,
		Color: t.Color,
	}
}

type TagDTO struct {
	ID    int64  `json:"id" validate:"omitempty,min=1"`
	Name  string `json:"name" validate:"required,min=1,max=50"`
	Color string `json:"color" validate:"omitempty,hexcolor"`
}

func (t *TagDTO) Validate() error {
	return validate.Struct(t)
}

func (t *TagDTO) ToModel() *TagModel {
	return &TagModel{
		ID:    t.ID,
		Name:  t.Name,
		Color: t.Color,
	}
}
