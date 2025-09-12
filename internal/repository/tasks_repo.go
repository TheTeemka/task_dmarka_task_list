package repository

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/TheTeemka/task_dmarka_task_list/internal/models"
	"github.com/TheTeemka/task_dmarka_task_list/pkg/merrors"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(task *models.TaskModel) (*models.TaskModel, error) {
	query := `
		INSERT INTO tasks (title, description, status, priority)
		VALUES(?, ?, ?, ?)
		RETURNING id`
	args := []any{task.Title, task.Description, task.Status, task.Priority}

	err := r.db.QueryRow(query, args...).Scan(&task.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (r *TaskRepo) GetTaskByID(id int64) (*models.TaskModel, error) {
	query := `
		SELECT id, title, description, status, priority
		FROM tasks
		WHERE id = ?`
	var task models.TaskModel

	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, merrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get task by id: %w", err)
	}

	return &task, nil
}

func (r *TaskRepo) UpdateTask(id int64, task *models.TaskModel) (*models.TaskModel, error) {
	query := `
		UPDATE tasks
		SET title = ?, description = ?, status = ?, priority = ?
		WHERE id = ?`

	args := []any{task.Title, task.Description, task.Status, task.Priority, id}

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return nil, merrors.ErrNotFound
	}

	return task, nil
}

func (r *TaskRepo) DeleteTask(id int64) error {
	query := `
		DELETE FROM tasks
		WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
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

func (r *TaskRepo) GetListByFilters(filter *models.TaskFilter) ([]*models.TaskModel, error) {
	b := squirrel.StatementBuilder.Select("id", "title", "description", "status", "priority").From("tasks")

	b = FilterToSQL(filter, b)
	query, args, err := b.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer rows.Close()

	var tasks []*models.TaskModel
	for rows.Next() {
		var task models.TaskModel
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tasks, nil
}

func FilterToSQL(f *models.TaskFilter, b squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f == nil {
		return b
	}

	if f.Limit != nil && *f.Limit > 0 {
		b = b.Limit(uint64(*f.Limit))
	}
	if f.Offset != nil && *f.Offset > 0 {
		b = b.Offset(uint64(*f.Offset))
	}

	if f.TagIDs != nil {
		b = b.Join("task_tags tt ON tasks.id = tt.task_id").
			Where(squirrel.Eq{"tt.tag_id": f.TagIDs})
	}

	if f.Priority != nil {
		b = b.Where(squirrel.Eq{"tasks.priority": *f.Priority})
	}

	if f.Status != nil {
		b = b.Where(squirrel.Eq{"tasks.status": *f.Status})
	}

	if f.SortBy != nil && *f.SortBy != "" {
		order := "ASC"
		if f.SortOrder != nil && *f.SortOrder == "desc" {
			order = "DESC"
		}
		b = b.OrderBy(fmt.Sprintf("%s %s", *f.SortBy, order))
	}

	return b
}
