package models

type CollectionVariable struct {
	ID           int    `json:"id" db:"id"`
	CollectionID int    `json:"collection_id" db:"collection_id"`
	Key          string `json:"key" db:"key"`
	Value        string `json:"value" db:"value"`
}
