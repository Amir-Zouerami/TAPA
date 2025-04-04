package models

import (
	"encoding/json"

	"github.com/Amir-Zouerami/TAPA/internal/types"
)

// AppState represents the last state of the application UI.
type AppState struct {
	ID                  int              `json:"id" db:"id"`
	SelectedEnvironment *int             `json:"selected_environment" db:"selected_environment"`
	SelectedTab         *int             `json:"selected_tab" db:"selected_tab"`
	FirstLaunch         bool             `json:"first_launch" db:"first_launch"`
	OpenTabs            []Tab            `json:"open_tabs" db:"-"` // Not directly mapped to DB (list of currently open tabs in the UI)
	OpenTabsJSON        string           `json:"-" db:"open_tabs"` // Raw JSON storage for OpenTabs (used only for database operations).
	UpdatedAt           types.SqliteTime `json:"updated_at" db:"updated_at"`
}

// Tab represents an open tab in the application UI.
type Tab struct {
	RequestID int  `json:"request_id"`
	IsSaved   bool `json:"is_saved"`
}

// BeforeSave serializes the OpenTabs slice to JSON for database storage.
func (a *AppState) BeforeSave() error {
	if a.OpenTabs == nil {
		a.OpenTabs = []Tab{}
	}

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
