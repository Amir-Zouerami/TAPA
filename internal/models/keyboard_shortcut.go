package models

type KeyboardShortcut struct {
	ID       int    `json:"id" db:"id"`
	Action   string `json:"action" db:"action"`
	Shortcut string `json:"shortcut" db:"shortcut"`
}
