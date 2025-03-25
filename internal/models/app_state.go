package models

import "time"

package models

import (
	"encoding/json"
	"time"
)

// AppState represents the current state of the application UI.
type AppState struct {
	// ID is the unique identifier for the application state record.
	ID int `json:"id" db:"id"`

	// SelectedEnvironment references the currently selected environment, if any.
	SelectedEnvironment *int `json:"selected_environment,omitempty" db:"selected_environment"`

	// OpenTabs contains the list of currently open tabs in the UI.
	OpenTabs []Tab `json:"open_tabs" db:"-"` // Not directly mapped to DB

	// Raw JSON storage for OpenTabs (used only for database operations).
	OpenTabsJSON string `json:"-" db:"open_tabs"`

	// UpdatedAt records when the application state was last modified.
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Tab represents an open tab in the application UI.
type Tab struct {
	// RequestID is a nullable reference to a request.
	RequestID *int `json:"request_id"`
	// IsSaved indicates whether there are unsaved changes.
	IsSaved bool `json:"is_saved"`
}

// BeforeSave serializes the OpenTabs slice to JSON for database storage.
func (a *AppState) BeforeSave() error {
	tabsJSON, err := json.Marshal(a.OpenTabs)
	if err != nil {
		return err
	}
	a.OpenTabsJSON = string(tabsJSON)
	return nil
}

// AfterLoad deserializes the OpenTabsJSON string into the OpenTabs slice.
func (a *AppState) AfterLoad() error {
	if a.OpenTabsJSON == "" {
		a.OpenTabs = []Tab{}
		return nil
	}
	return json.Unmarshal([]byte(a.OpenTabsJSON), &a.OpenTabs)
}
