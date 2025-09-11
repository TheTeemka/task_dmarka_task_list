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

// GetTagsForTask returns full tag records associated with a task.
func (r *TagsRepo) GetTagsForTask(taskID int64) ([]models.Tag, error) {
	query := `
		SELECT t.id, t.name, t.color
		FROM tags t
		JOIN task_tags tt ON t.id = tt.tag_id
		WHERE tt.task_id = ?`

	var tags []models.Tag
	rows, err := r.db.Query(query, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags for task: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tag models.Tag
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

func (r *TagsRepo) GetTagByID(id int64) (*models.Tag, error) {
	query := `
		SELECT id, name, color
		FROM tags
		WHERE id = ?`
	var tag models.Tag

	err := r.db.QueryRow(query, id).Scan(&tag.ID, &tag.Name, &tag.Color)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, merrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get tag by id: %w", err)
	}

	return &tag, nil
}
