package repository

import (
	"context"

	"github.com/Amir-Zouerami/TAPA/internal/models"
	"github.com/jmoiron/sqlx"
)

type CollectionsRepositoryInterface interface {
	GetAllCollections(ctx context.Context) ([]models.Collection, error)
	GetAllFolders(ctx context.Context) ([]models.Folder, error)
	GetAllRequestSummaries(ctx context.Context) ([]models.RequestSummary, error)
}

type CollectionsRepository struct {
	db *sqlx.DB
}

// GetAllCollections returns all collections from the database.
func (r *CollectionsRepository) GetAllCollections(ctx context.Context) ([]models.Collection, error) {
	var cols []models.Collection
	query := `
		SELECT id, name, description, position, updated_at
		FROM collections
		ORDER BY position ASC`

	if err := r.db.SelectContext(ctx, &cols, query); err != nil {
		return nil, err
	}

	return cols, nil
}

// GetAllFolders returns all folders from the database.
func (r *CollectionsRepository) GetAllFolders(ctx context.Context) ([]models.Folder, error) {
	var folders []models.Folder
	query := `
		SELECT id, collection_id, name, position, updated_at 
		FROM folders 
		ORDER BY position ASC`

	if err := r.db.SelectContext(ctx, &folders, query); err != nil {
		return nil, err
	}

	return folders, nil
}

// GetAllRequestSummaries returns minimal request data: id, collection_id, folder_id, name, and method.
func (r *CollectionsRepository) GetAllRequestSummaries(ctx context.Context) ([]models.RequestSummary, error) {
	var reqs []models.RequestSummary
	query := `
		SELECT id, collection_id, folder_id, name, method 
		FROM requests`

	if err := r.db.SelectContext(ctx, &reqs, query); err != nil {
		return nil, err
	}

	return reqs, nil
}

func NewCollectionsRepository(db *sqlx.DB) *CollectionsRepository {
	return &CollectionsRepository{db: db}
}
