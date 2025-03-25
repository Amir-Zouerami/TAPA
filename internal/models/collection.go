package models

import "time"

type Collection struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description,omitempty" db:"description"`
	Position    int       `json:"position" db:"position"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type PopulatedCollection struct {
	Collection Collection        `json:"collection"`
	Folders    []PopulatedFolder `json:"folders"`
	Requests   []RequestBasic    `json:"requests"` // requests with no folder.
}
