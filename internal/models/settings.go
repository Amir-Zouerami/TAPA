package models

type UserSettings struct {
	ID         int    `json:"id" db:"id"`
	Theme      string `json:"theme" db:"theme"`             // Expected: "light", "dark", "system"
	MaxHistory int    `json:"max_history" db:"max_history"` // Maximum history records allowed
	Language   string `json:"language" db:"language"`
	FontFamily string `json:"font_family" db:"font_family"`
	FontSize   int    `json:"font_size" db:"font_size"`
}
