package models

type RequestCookie struct {
	ID        int    `json:"id" db:"id"`
	RequestID int    `json:"request_id" db:"request_id"`
	Key       string `json:"key" db:"key"`
	Value     string `json:"value,omitempty" db:"value"`
}
