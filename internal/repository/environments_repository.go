package repository

import (
	"context"

	"github.com/Amir-Zouerami/TAPA/internal/models"
	"github.com/jmoiron/sqlx"
)

type EnvironmentsRepositoryInterface interface {
	ListAllEnvironments(ctx context.Context) ([]models.Environment, error)
	GetEnvironmentDetails(ctx context.Context, envID int) ([]models.EnvironmentVariable, error)
}

type EnvironmentsRepository struct {
	db *sqlx.DB
}

func (r *EnvironmentsRepository) ListAllEnvironments(ctx context.Context) ([]models.Environment, error) {
	var environmentsList []models.Environment
	query := `
	SELECT id, name, updated_at FROM environments
	`

	if err := r.db.SelectContext(ctx, &environmentsList, query); err != nil {
		return nil, err
	}

	return environmentsList, nil
}

func (r *EnvironmentsRepository) GetEnvironmentDetails(ctx context.Context, envID int) ([]models.EnvironmentVariable, error) {
	var allVariables []models.EnvironmentVariable

	err := r.db.Select(
		&allVariables,
		"SELECT key, value FROM environment_variables WHERE environment_id = ?",
		envID,
	)

	if err != nil {
		return nil, err
	}

	return allVariables, nil
}

func NewEnvironmentsRepository(db *sqlx.DB) *EnvironmentsRepository {
	return &EnvironmentsRepository{db: db}
}
