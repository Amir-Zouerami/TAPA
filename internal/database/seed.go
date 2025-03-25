package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// PlaceholderPost represents a sample post fetched from JSONPlaceholder.
type PlaceholderPost struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// SeedDB inserts realistic sample data into all tables.
func SeedDB(db *sqlx.DB) error {
	log.Println("Seeding database...")

	if err := seedCollection(db); err != nil {
		return err
	}

	if err := seedFolders(db); err != nil {
		return err
	}

	if err := seedCollectionVariables(db); err != nil {
		return err
	}

	if err := seedRequests(db); err != nil {
		return err
	}

	if err := seedEnvironment(db); err != nil {
		return err
	}

	if err := seedSyncMetadata(db); err != nil {
		return err
	}

	if err := seedKeyboardShortcuts(db); err != nil {
		return err
	}

	if err := seedUserSettings(db); err != nil {
		return err
	}

	log.Println("Database seeded successfully.")
	return nil
}

// seedCollection inserts a sample collection with a distinguishable ID.
func seedCollection(db *sqlx.DB) error {
	_, err := db.Exec(`
		INSERT INTO collections (id, name, description, position) VALUES
		(101, 'Seeded Collection', 'A sample collection to store API requests and related data', 1);
	`)
	if err != nil {
		return fmt.Errorf("failed to insert collection: %w", err)
	}
	return nil
}

// seedFolders inserts a sample folder associated with the seeded collection.
func seedFolders(db *sqlx.DB) error {
	_, err := db.Exec(`
		INSERT INTO folders (id, collection_id, name, position) VALUES
		(201, 101, 'Seeded General Requests', 1);
	`)
	if err != nil {
		return fmt.Errorf("failed to insert folder: %w", err)
	}
	return nil
}

// seedCollectionVariables inserts sample variables for the seeded collection.
func seedCollectionVariables(db *sqlx.DB) error {
	_, err := db.Exec(`
		INSERT INTO collection_variables (id, collection_id, key, value) VALUES
		(301, 101, 'defaultTimeout', '30000'),
		(302, 101, 'retryAttempts', '3');
	`)
	if err != nil {
		return fmt.Errorf("failed to insert collection variables: %w", err)
	}
	return nil
}

// seedRequests inserts sample requests and their dependencies.
func seedRequests(db *sqlx.DB) error {
	posts := []PlaceholderPost{
		{
			ID:    1,
			Title: "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
			Body:  "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
		},
		{
			ID:    2,
			Title: "qui est esse",
			Body:  "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
		},
	}

	log.Println("Seeding requests and associated data...")

	for _, post := range posts {
		if err := seedRequestAndDependencies(db, post); err != nil {
			return err
		}
	}
	return nil
}

// seedRequestAndDependencies inserts one request and its associated data.
// We offset the request ID by 1000 (e.g., 1 becomes 1001) and use a similar value for the position.
func seedRequestAndDependencies(db *sqlx.DB, post PlaceholderPost) error {
	newID := post.ID + 1000

	_, err := db.Exec(`
		INSERT INTO requests 
		(id, collection_id, folder_id, position, name, method, url, body, timeout, allow_redirects, ssl_verification, remove_referer_on_redirect, encode_url)
		VALUES (?, 101, 201, ?, ?, 'GET', ?, '', 30000, 1, 1, 0, 1);
	`, newID, newID, post.Title, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", post.ID))

	if err != nil {
		return fmt.Errorf("failed to insert request '%s': %w", post.Title, err)
	}

	if err := seedRequestHeaders(db, newID, post); err != nil {
		return err
	}

	if err := seedRequestQueryParams(db, newID, post); err != nil {
		return err
	}

	if err := seedRequestHistory(db, newID, post); err != nil {
		return err
	}

	if err := seedRequestScripts(db, newID, post); err != nil {
		return err
	}

	return nil
}

// seedRequestHeaders inserts a sample header for a given request.
// Note: The id field is omitted to let the database autogenerate it.
func seedRequestHeaders(db *sqlx.DB, requestID int, post PlaceholderPost) error {
	_, err := db.Exec(`
		INSERT INTO request_headers (request_id, key, value) VALUES 
		(?, 'Content-Type', 'application/json');
	`, requestID)
	if err != nil {
		return fmt.Errorf("failed to insert header for request '%s': %w", post.Title, err)
	}
	return nil
}

// seedRequestQueryParams inserts a sample query parameter for a given request.
// The id field is omitted.
func seedRequestQueryParams(db *sqlx.DB, requestID int, post PlaceholderPost) error {
	_, err := db.Exec(`
		INSERT INTO request_query_params (request_id, key, value) VALUES 
		(?, 'userId', '1');
	`, requestID)
	if err != nil {
		return fmt.Errorf("failed to insert query parameter for request '%s': %w", post.Title, err)
	}
	return nil
}

// seedRequestHistory inserts a sample history record for a given request.
func seedRequestHistory(db *sqlx.DB, requestID int, post PlaceholderPost) error {
	now := time.Now().Format(time.DateTime)
	_, err := db.Exec(`
		INSERT INTO request_history 
		(request_id, timestamp, method, url, headers, query_params, body, status_code, response_time, data_volume)
		VALUES (?, ?, 'GET', ?, '{"Content-Type":"application/json"}', '{"userId":"1"}', '', 200, 123, 456);
	`, requestID, now, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", post.ID))
	if err != nil {
		return fmt.Errorf("failed to insert history for request '%s': %w", post.Title, err)
	}
	return nil
}

// seedRequestScripts inserts two sample scripts for a given request.
// The id field is omitted so that autoincrement works.
func seedRequestScripts(db *sqlx.DB, requestID int, post PlaceholderPost) error {
	_, err := db.Exec(`
		INSERT INTO request_scripts (request_id, script) VALUES 
		(?, ?);
	`, requestID, fmt.Sprintf("console.log('Pre-request script for post %d');", post.ID))
	if err != nil {
		return fmt.Errorf("failed to insert pre-request script for request '%s': %w", post.Title, err)
	}

	_, err = db.Exec(`
		INSERT INTO request_scripts (request_id, script) VALUES 
		(?, ?);
	`, requestID, fmt.Sprintf("console.log('Test script for post %d');", post.ID))
	if err != nil {
		return fmt.Errorf("failed to insert test script for request '%s': %w", post.Title, err)
	}
	return nil
}

// seedEnvironment inserts a sample environment and its variables.
func seedEnvironment(db *sqlx.DB) error {
	// ID 401 for environment.
	_, err := db.Exec(`
		INSERT INTO environments (id, name) VALUES
		(401, 'Seeded Development');
	`)
	if err != nil {
		return fmt.Errorf("failed to insert environment: %w", err)
	}

	_, err = db.Exec(`
		INSERT INTO environment_variables (id, environment_id, key, value) VALUES
		(501, 401, 'BASE_URL', 'https://jsonplaceholder.typicode.com'),
		(502, 401, 'HOST', 'jsonplaceholder.typicode.com'),
		(503, 401, 'PORT', '443'),
		(504, 401, 'PATH', '/posts');
	`)
	if err != nil {
		return fmt.Errorf("failed to insert environment variables: %w", err)
	}
	return nil
}

// seedSyncMetadata inserts sample synchronization metadata.
func seedSyncMetadata(db *sqlx.DB) error {
	_, err := db.Exec(`
		INSERT INTO sync_metadata (id, entity_type, entity_id, last_updated, is_dirty) VALUES
		(601, 'collections', 101, datetime('now'), 0);
	`)
	if err != nil {
		return fmt.Errorf("failed to insert sync metadata: %w", err)
	}
	return nil
}

// seedKeyboardShortcuts inserts sample keyboard shortcuts.
func seedKeyboardShortcuts(db *sqlx.DB) error {
	_, err := db.Exec(`
		INSERT INTO keyboard_shortcuts (id, action, shortcut) VALUES
		(701, 'Save Request', 'Ctrl+S'),
		(702, 'Send Request', 'Ctrl+Enter');
	`)
	if err != nil {
		return fmt.Errorf("failed to insert keyboard shortcuts: %w", err)
	}
	return nil
}

// seedUserSettings inserts a sample user settings record.
func seedUserSettings(db *sqlx.DB) error {
	_, err := db.Exec(`
		INSERT INTO user_settings (id, theme, max_history, language, font_family, font_size) VALUES
		(1, 'system', 200, 'en', 'Arial', 14);
	`)
	if err != nil {
		return fmt.Errorf("failed to insert user settings: %w", err)
	}
	return nil
}

// IsSeeded checks if the seed data exists by looking for a known record.
func IsSeeded(db *sqlx.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(1) FROM collections WHERE id = 101").Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check seed status: %w", err)
	}
	return count > 0, nil
}

// FlushDB clears all tables.
func FlushDB(db *sqlx.DB) error {
	log.Println("Flushing database...")

	tables := []string{
		"collections", "folders", "requests", "request_headers", "request_query_params",
		"request_cookies", "environments", "environment_variables", "collection_variables",
		"request_history", "request_scripts", "sync_metadata", "keyboard_shortcuts", "user_settings", "app_state",
	}

	_, _ = db.Exec("PRAGMA foreign_keys = OFF;")
	for _, table := range tables {
		_, err := db.Exec("DELETE FROM " + table)
		if err != nil {
			_, _ = db.Exec("PRAGMA foreign_keys = ON;")
			return fmt.Errorf("failed to flush table %s: %w", table, err)
		}
	}
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")
	log.Println("Database flushed successfully.")
	return nil
}
