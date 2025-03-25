package models

import "time"

type TestResult struct {
	ID        int       `json:"id" db:"id"`
	RequestID int       `json:"request_id" db:"request_id"`
	TestName  string    `json:"test_name" db:"test_name"`
	Result    string    `json:"result,omitempty" db:"result"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
