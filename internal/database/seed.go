package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PlaceholderPost struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// SeedDB inserts sample data into the database.
func SeedDB(db *sql.DB) error {
	log.Println("Seeding database...")

	// Collections
	_, err := db.Exec(`INSERT INTO collections (id, name, description) VALUES
		('col-1', 'Sample Collection', 'A test collection');`)
	if err != nil {
		return err
	}

	// Fetch JSONPlaceholder posts as sample requests
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?_limit=2")
	if err != nil {
		return fmt.Errorf("failed to fetch posts: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var posts []PlaceholderPost
	if err := json.Unmarshal(body, &posts); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	log.Println("Seeding requests...")
	for _, post := range posts {
		_, err := db.Exec(`INSERT INTO requests (id, collection_id, name, method, url, body, timeout) VALUES
			(?, 'col-1', ?, 'GET', 'https://jsonplaceholder.typicode.com/posts/' || ?, '', 30000);`,
			fmt.Sprintf("req-%d", post.ID), post.Title, post.ID)
		if err != nil {
			return fmt.Errorf("failed to insert request %s: %w", post.Title, err)
		}
	}

	// Environments
	_, err = db.Exec(`INSERT INTO environments (id, name) VALUES
		('env-1', 'Development');`)
	if err != nil {
		return err
	}

	// Environment Variables
	_, err = db.Exec(`INSERT INTO environment_variables (id, environment_id, key, value) VALUES
		('envvar-1', 'env-1', 'BASE_URL', 'https://jsonplaceholder.typicode.com');`)
	if err != nil {
		return err
	}

	log.Println("Database seeded successfully.")
	return nil
}

// FlushDB clears all tables.
func FlushDB(db *sql.DB) error {
	log.Println("Flushing database...")
	tables := []string{
		"collections", "folders", "requests", "request_headers", "request_query_params",
		"request_auth", "environments", "environment_variables", "collection_variables",
		"request_history", "request_scripts", "sync_metadata", "keyboard_shortcuts", "user_settings",
	}

	for _, table := range tables {
		_, err := db.Exec("DELETE FROM " + table)
		if err != nil {
			return err
		}
	}

	log.Println("Database flushed successfully.")
	return nil
}
