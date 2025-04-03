package models

import (
	"github.com/Amir-Zouerami/TAPA/internal/types"
)

type Environment struct {
	ID        int              `json:"id" db:"id"`
	Name      string           `json:"name" db:"name"`
	UpdatedAt types.SqliteTime `json:"updated_at" db:"updated_at"`
}
