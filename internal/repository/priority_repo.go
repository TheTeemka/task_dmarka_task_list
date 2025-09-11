package repository

import (
	"database/sql"
	"fmt"

	"github.com/TheTeemka/task_dmarka_task_list/internal/models"
	"github.com/TheTeemka/task_dmarka_task_list/pkg/merrors"
)

type PriorityRepo struct {
	db *sql.DB
}

func NewPriorityRepo(db *sql.DB) *PriorityRepo {
	return &PriorityRepo{db: db}
}

func (r *PriorityRepo) GetByID(id int64) (*models.Priority, error) {
	var p models.Priority
	query := `SELECT id, name, color FROM priorities WHERE id = ?`
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&p.ID, &p.Name, &p.Color); err != nil {
		if err == sql.ErrNoRows {
			return nil, merrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get priority by id: %w", err)
	}
	return &p, nil
}

func (r *PriorityRepo) Create(p *models.Priority) (*models.Priority, error) {
	query := `
		INSERT INTO priorities (name, color)
		VALUES (?, ?)
		RETURNING id`
	if err := r.db.QueryRow(query, p.Name, p.Color).Scan(&p.ID); err != nil {
		return nil, fmt.Errorf("failed to create priority: %w", err)
	}
	return p, nil
}

func (r *PriorityRepo) Update(p *models.Priority) (*models.Priority, error) {
	query := `
		UPDATE priorities
		SET name = ?, color = ?
		WHERE id = ?`
	res, err := r.db.Exec(query, p.Name, p.Color, p.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update priority: %w", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, merrors.ErrNotFound
	}
	return p, nil
}

func (r *PriorityRepo) Delete(id int64) error {
	query := `DELETE FROM priorities WHERE id = ?`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete priority: %w", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return merrors.ErrNotFound
	}
	return nil
}
