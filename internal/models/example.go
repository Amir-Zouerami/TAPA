package models

import "time"

type RequestExample struct {
	ID              int       `json:"id" db:"id"`
	RequestID       int       `json:"request_id" db:"request_id"`
	Timestamp       time.Time `json:"timestamp" db:"timestamp"`
	Method          string    `json:"method" db:"method"`
	URL             string    `json:"url" db:"url"`
	Headers         string    `json:"headers,omitempty" db:"headers"`
	QueryParams     string    `json:"query_params,omitempty" db:"query_params"`
	Body            string    `json:"body,omitempty" db:"body"`
	StatusCode      int       `json:"status_code" db:"status_code"`
	Response        string    `json:"response,omitempty" db:"response"`
	ResponseHeaders string    `json:"response_headers,omitempty" db:"response_headers"`
	ResponseCookies string    `json:"response_cookies,omitempty" db:"response_cookies"`
	ResponseTime    int       `json:"response_time" db:"response_time"`
	DataVolume      int       `json:"data_volume" db:"data_volume"`
}
