package models

import "time"

type SyncMetadata struct {
	ID          int       `json:"id" db:"id"`
	EntityType  string    `json:"entity_type" db:"entity_type"` // 'requests', 'collections', 'variables', 'history'
	EntityID    int       `json:"entity_id" db:"entity_id"`
	LastUpdated time.Time `json:"last_updated" db:"last_updated"`
	IsDirty     bool      `json:"is_dirty" db:"is_dirty"`
}
