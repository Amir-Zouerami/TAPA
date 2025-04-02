package models

import (
	"github.com/Amir-Zouerami/TAPA/internal/types"
)

type Folder struct {
	ID           int              `json:"id" db:"id"`
	CollectionID int              `json:"collection_id" db:"collection_id"`
	Name         string           `json:"name" db:"name"`
	Position     int              `json:"position" db:"position"`
	UpdatedAt    types.SqliteTime `json:"updated_at" db:"updated_at"`
}

type PopulatedFolder struct {
	Folder   Folder           `json:"folder"`
	Requests []RequestSummary `json:"requests"`
}
