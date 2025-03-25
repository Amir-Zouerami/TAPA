package models

import "time"

type Folder struct {
	ID           int       `json:"id" db:"id"`
	CollectionID int       `json:"collection_id" db:"collection_id"`
	Name         string    `json:"name" db:"name"`
	Position     int       `json:"position" db:"position"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type PopulatedFolder struct {
	Folder   Folder         `json:"folder"`
	Requests []RequestBasic `json:"requests"`
}
