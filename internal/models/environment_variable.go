package models

type FullEnvironmentVariable struct {
	ID            int    `json:"id" db:"id"`
	EnvironmentID int    `json:"environment_id" db:"environment_id"`
	Key           string `json:"key" db:"key"`
	Value         string `json:"value" db:"value"`
}

type EnvironmentVariable struct {
	Key   string `json:"key" db:"key"`
	Value string `json:"value" db:"value"`
}
