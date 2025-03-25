package models

import "time"

type Request struct {
	ID                      int       `json:"id" db:"id"`
	CollectionID            *int      `json:"collection_id,omitempty" db:"collection_id"`
	FolderID                *int      `json:"folder_id,omitempty" db:"folder_id"`
	Position                int       `json:"position" db:"position"`
	Name                    string    `json:"name" db:"name"`
	Method                  string    `json:"method" db:"method"`
	URL                     string    `json:"url" db:"url"`
	Body                    string    `json:"body,omitempty" db:"body"`
	BodyFormat              string    `json:"body_format" db:"body_format"` // "JSON", "XML", "form-data", "raw"
	Notes                   string    `json:"notes,omitempty" db:"notes"`
	Timeout                 int       `json:"timeout" db:"timeout"`
	AllowRedirects          bool      `json:"allow_redirects" db:"allow_redirects"`
	SSLVerification         bool      `json:"ssl_verification" db:"ssl_verification"`
	RemoveRefererOnRedirect bool      `json:"remove_referer_on_redirect" db:"remove_referer_on_redirect"`
	EncodeURL               bool      `json:"encode_url" db:"encode_url"`
	CreatedAt               time.Time `json:"created_at" db:"created_at"`
	UpdatedAt               time.Time `json:"updated_at" db:"updated_at"`
}

type RequestBasic struct {
	ID           int    `json:"id" db:"id"`
	CollectionID *int   `json:"collection_id,omitempty" db:"collection_id"`
	FolderID     *int   `json:"folder_id,omitempty" db:"folder_id"`
	Name         string `json:"name" db:"name"`
	Method       string `json:"method" db:"method"`
}

type FullDashboardRequestList struct {
	Collections   []PopulatedCollection `json:"collections"`
	LooseRequests []RequestBasic        `json:"loose_requests"`
}

// type RequestWithDetail struct {
// 	Request
// 	Headers     []RequestHeader     `json:"headers"`
// 	QueryParams []RequestQueryParam `json:"query_params"`
// 	Cookies     []RequestCookie     `json:"cookies"`
// 	Scripts     []RequestScript     `json:"scripts"`
// }
