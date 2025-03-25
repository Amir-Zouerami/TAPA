package models

import "time"

type AppState struct {
	ID                  int       `json:"id" db:"id"`
	SelectedEnvironment *int      `json:"selected_environment,omitempty" db:"selected_environment"`
	OpenTabs            string    `json:"open_tabs" db:"open_tabs"` // JSON string representing an array of open tab objects, e.g. [{"request_id":12, "is_saved":false}, {"request_id":null, "is_saved":false}]
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}
