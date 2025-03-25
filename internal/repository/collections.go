package repository

import (
	"github.com/Amir-Zouerami/TAPA/internal/errors"
	"github.com/Amir-Zouerami/TAPA/internal/models"
	"github.com/jmoiron/sqlx"
)

type CollectionsRepository struct {
	db *sqlx.DB
}

// GetAllCollections returns all collections from the database.
func (r *CollectionsRepository) GetAllCollections() ([]models.Collection, error) {
	var cols []models.Collection
	query := `
		SELECT id, name, description, position, created_at, updated_at
		FROM collections
		ORDER BY position ASC`

	if err := r.db.Select(&cols, query); err != nil {
		return nil, errors.Wrap(errors.ErrCollectionsRetrieval, err)
	}

	return cols, nil
}

// GetAllFolders returns all folders from the database.
func (r *CollectionsRepository) GetAllFolders() ([]models.Folder, error) {
	var folders []models.Folder
	query := `
		SELECT id, collection_id, name, position, created_at, updated_at 
		FROM folders 
		ORDER BY position ASC`

	if err := r.db.Select(&folders, query); err != nil {
		return nil, errors.Wrap(errors.ErrFoldersRetrieval, err)
	}

	return folders, nil
}

// GetAllRequestSummaries returns minimal request data: id, collection_id, folder_id, name, and method.
func (r *CollectionsRepository) GetAllRequestSummaries() ([]models.RequestBasic, error) {
	var reqs []models.RequestBasic
	query := `
		SELECT id, collection_id, folder_id, name, method 
		FROM requests`

	if err := r.db.Select(&reqs, query); err != nil {
		return nil, errors.Wrap(errors.ErrRequestSummariesRetrieval, err)
	}

	return reqs, nil
}

func NewCollectionsRepository(db *sqlx.DB) *CollectionsRepository {
	return &CollectionsRepository{db: db}
}
