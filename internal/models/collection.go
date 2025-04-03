package models

import (
	"github.com/Amir-Zouerami/TAPA/internal/types"
)

type Collection struct {
	ID          int              `json:"id" db:"id"`
	Name        string           `json:"name" db:"name"`
	Description string           `json:"description,omitempty" db:"description"`
	Position    int              `json:"position" db:"position"`
	UpdatedAt   types.SqliteTime `json:"updated_at" db:"updated_at"`
}

// type PopulatedCollection struct {
// 	Collection Collection        `json:"collection"`
// 	Folders    []PopulatedFolder `json:"folders"`
// 	Requests   []RequestSummary    `json:"requests"` // requests with no folder.
// }
