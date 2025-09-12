package repository

import (
	"database/sql"
	"fmt"

	"github.com/TheTeemka/task_dmarka_task_list/internal/models"
	"github.com/TheTeemka/task_dmarka_task_list/pkg/merrors"
)

type TagsRepo struct {
	db *sql.DB
}

func NewTagsRepo(db *sql.DB) *TagsRepo {
	return &TagsRepo{db: db}
}

func (r *TagsRepo) CreateTag(tag *models.TagModel) (*models.TagModel, error) {
	query := `INSERT INTO tags (name, color) VALUES (?, ?)`
	result, err := r.db.Exec(query, tag.Name, tag.Color)
	if err != nil {
		return nil, fmt.Errorf("failed to create tag: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %w", err)
	}
	tag.ID = id
	return tag, nil
}

// GetTagsForTask returns full tag records associated with a task.
func (r *TagsRepo) GetTagsForTask(taskID int64) ([]models.TagModel, error) {
	query := `
		SELECT t.id, t.name, t.color
		FROM tags t
		JOIN task_tags tt ON t.id = tt.tag_id
		WHERE tt.task_id = ?`

	var tags []models.TagModel
	rows, err := r.db.Query(query, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags for task: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tag models.TagModel
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Color); err != nil {
			return nil, fmt.Errorf("failed to scan tag: %w", err)
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tags, nil
}

func (r *TagsRepo) GetTagByID(id int64) (*models.TagModel, error) {
	query := `
		SELECT id, name, color
		FROM tags
		WHERE id = ?`
	var tag models.TagModel

	err := r.db.QueryRow(query, id).Scan(&tag.ID, &tag.Name, &tag.Color)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, merrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get tag by id: %w", err)
	}

	return &tag, nil
}

func (r *TagsRepo) UpdateTag(tag *models.TagModel) (*models.TagModel, error) {
	query := `UPDATE tags SET name = ?, color = ? WHERE id = ?`
	_, err := r.db.Exec(query, tag.Name, tag.Color, tag.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update tag: %w", err)
	}
	return tag, nil
}

func (r *TagsRepo) DeleteTag(id int64) error {
	query := `DELETE FROM tags WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return merrors.ErrNotFound
	}

	return nil
}

func (r *TagsRepo) GetAllTags() ([]*models.TagModel, error) {
	query := `SELECT id, name, color FROM tags ORDER BY name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tags: %w", err)
	}
	defer rows.Close()

	var tags []*models.TagModel
	for rows.Next() {
		var tag models.TagModel
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Color); err != nil {
			return nil, fmt.Errorf("failed to scan tag: %w", err)
		}
		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tags, nil
}

func (r *TagsRepo) AddTagToTask(taskID, tagID int64) error {
	query := `INSERT INTO task_tags (task_id, tag_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, taskID, tagID)
	if err != nil {
		return fmt.Errorf("failed to add tag to task: %w", err)
	}
	return nil
}

func (r *TagsRepo) RemoveTagFromTask(taskID, tagID int64) error {
	query := `DELETE FROM task_tags WHERE task_id = ? AND tag_id = ?`
	result, err := r.db.Exec(query, taskID, tagID)
	if err != nil {
		return fmt.Errorf("failed to remove tag from task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return merrors.ErrNotFound
	}

	return nil
}
