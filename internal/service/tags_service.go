package service

import (
	"github.com/TheTeemka/task_dmarka_task_list/internal/models"
	"github.com/TheTeemka/task_dmarka_task_list/internal/repository"
)

type TagsService struct {
	tagsRepo *repository.TagsRepo
}

func NewTagsService(tagRepo *repository.TagsRepo) *TagsService {
	return &TagsService{tagsRepo: tagRepo}
}

func (ts *TagsService) CreateNewTag(tag *models.TagDTO) (*models.TagDTO, error) {
	if err := tag.Validate(); err != nil {
		return nil, err
	}

	createdTag, err := ts.tagsRepo.CreateTag(tag.ToModel())
	if err != nil {
		return nil, err
	}

	return createdTag.ToDTO(), nil
}

func (ts *TagsService) GetTagByID(id int64) (*models.TagDTO, error) {
	tagModel, err := ts.tagsRepo.GetTagByID(id)
	if err != nil {
		return nil, err
	}

	return tagModel.ToDTO(), nil
}

func (ts *TagsService) UpdateTag(tag *models.TagDTO) (*models.TagDTO, error) {
	if err := tag.Validate(); err != nil {
		return nil, err
	}

	updatedTag, err := ts.tagsRepo.UpdateTag(tag.ToModel())
	if err != nil {
		return nil, err
	}

	return updatedTag.ToDTO(), nil
}

func (ts *TagsService) DeleteTag(id int64) error {
	return ts.tagsRepo.DeleteTag(id)
}

func (ts *TagsService) GetAllTags() ([]*models.TagDTO, error) {
	tagModels, err := ts.tagsRepo.GetAllTags()
	if err != nil {
		return nil, err
	}

	var tags []*models.TagDTO
	for _, tm := range tagModels {
		tags = append(tags, tm.ToDTO())
	}
	return tags, nil
}

func (ts *TagsService) GetTagsForTask(taskID int64) ([]*models.TagDTO, error) {
	tagModels, err := ts.tagsRepo.GetTagsForTask(taskID)
	if err != nil {
		return nil, err
	}

	var tags []*models.TagDTO
	for _, tm := range tagModels {
		tags = append(tags, tm.ToDTO())
	}
	return tags, nil
}

func (ts *TagsService) AddTagToTask(taskID, tagID int64) error {
	return ts.tagsRepo.AddTagToTask(taskID, tagID)
}

func (ts *TagsService) RemoveTagFromTask(taskID, tagID int64) error {
	return ts.tagsRepo.RemoveTagFromTask(taskID, tagID)
}
