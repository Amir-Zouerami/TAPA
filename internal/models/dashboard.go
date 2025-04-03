package models

type DashboardData struct {
	Environments                []Environment            `json:"environments"`
	CurrentEnvironmentVariables map[string]string        `json:"current_env_variables"`
	KeyboardShortcuts           map[int]KeyboardShortcut `json:"keyboard_shortcuts"`
	UserSettings                UserSettings             `json:"user_settings"`
	AppState                    AppState                 `json:"app_state"`
}
