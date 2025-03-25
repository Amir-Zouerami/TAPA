package models

type RequestScript struct {
	ID        int    `json:"id" db:"id"`
	RequestID int    `json:"request_id" db:"request_id"`
	Script    string `json:"script" db:"script"`
}
