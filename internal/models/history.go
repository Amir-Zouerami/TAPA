package models

import "time"

type RequestHistory struct {
	ID           int       `json:"id" db:"id"`
	RequestID    int       `json:"request_id" db:"request_id"`
	Timestamp    time.Time `json:"timestamp" db:"timestamp"`
	Method       string    `json:"method" db:"method"`
	URL          string    `json:"url" db:"url"`
	Headers      string    `json:"headers,omitempty" db:"headers"`           // Stored as JSON string
	QueryParams  string    `json:"query_params,omitempty" db:"query_params"` // Stored as JSON string
	Body         string    `json:"body,omitempty" db:"body"`
	StatusCode   int       `json:"status_code" db:"status_code"`
	ResponseTime int       `json:"response_time" db:"response_time"`
	DataVolume   int       `json:"data_volume" db:"data_volume"`
}
