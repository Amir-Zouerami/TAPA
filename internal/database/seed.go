package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// PlaceholderPost represents a sample post fetched from JSONPlaceholder.
type PlaceholderPost struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// SeedDB inserts realistic sample data into all tables.
func SeedDB(db *sql.DB) error {
	log.Println("Seeding database...")

	// Insert a sample collection.
	_, err := db.Exec(`INSERT INTO collections (id, name, description) VALUES
		('TAPA-DEV-COL-1', 'Sample Collection', 'A collection to store API requests and related data');`)
	if err != nil {
		return fmt.Errorf("failed to insert collection: %w", err)
	}

	// Insert a sample folder for the collection.
	_, err = db.Exec(`INSERT INTO folders (id, collection_id, name, position) VALUES
		('TAPA-DEV-FOLDER-1', 'TAPA-DEV-COL-1', 'General Requests', 1);`)
	if err != nil {
		return fmt.Errorf("failed to insert folder: %w", err)
	}

	// Fetch sample posts from JSONPlaceholder.
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

	log.Println("Seeding requests and associated data...")
	for _, post := range posts {
		// Use a consistent request ID.
		requestID := fmt.Sprintf("req-%d", post.ID)
		// Insert the request record into the folder and collection.
		_, err := db.Exec(`INSERT INTO requests 
			(id, collection_id, folder_id, name, method, url, body, timeout, allow_redirects, ssl_verification, remove_referer_on_redirect, encode_url)
			VALUES (?, 'TAPA-DEV-COL-1', 'TAPA-DEV-FOLDER-1', ?, 'GET', ?, '', 30000, 1, 1, 0, 1);`,
			requestID, post.Title, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", post.ID))
		if err != nil {
			return fmt.Errorf("failed to insert request %s: %w", post.Title, err)
		}

		// Insert a sample header for the request.
		headerID := fmt.Sprintf("hdr-%d", post.ID)
		_, err = db.Exec(`INSERT INTO request_headers (id, request_id, key, value) VALUES 
			(?, ?, 'Content-Type', 'application/json');`, headerID, requestID)
		if err != nil {
			return fmt.Errorf("failed to insert header for request %s: %w", post.Title, err)
		}

		// Insert a sample query parameter.
		queryParamID := fmt.Sprintf("qry-%d", post.ID)
		_, err = db.Exec(`INSERT INTO request_query_params (id, request_id, key, value) VALUES 
			(?, ?, 'userId', '1');`, queryParamID, requestID)
		if err != nil {
			return fmt.Errorf("failed to insert query parameter for request %s: %w", post.Title, err)
		}

		// Insert a sample auth record for the first request.
		if post.ID == 1 {
			_, err = db.Exec(`INSERT INTO request_auth (id, request_id, auth_type, token, username, password) VALUES 
				('auth-1', ?, 'Bearer', 'sample-token-123', '', '');`, requestID)
			if err != nil {
				return fmt.Errorf("failed to insert auth for request %s: %w", post.Title, err)
			}
		}

		// Insert a sample request history entry.
		historyID := fmt.Sprintf("hist-%d", post.ID)
		now := time.Now().Format(time.DateTime)
		_, err = db.Exec(`INSERT INTO request_history 
			(id, request_id, timestamp, history_type, method, url, headers, query_params, auth, body, status_code, response, response_headers)
			VALUES (?, ?, ?, 'auto', 'GET', ?, '{"Content-Type":"application/json"}', '{"userId":"1"}', '', '', 200, '{"id":1}', '{"Content-Length":"123"}');`,
			historyID, requestID, now, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", post.ID))
		if err != nil {
			return fmt.Errorf("failed to insert history for request %s: %w", post.Title, err)
		}

		// Insert sample request scripts: pre-request and test.
		preScriptID := fmt.Sprintf("script-pre-%d", post.ID)
		testScriptID := fmt.Sprintf("script-test-%d", post.ID)
		_, err = db.Exec(`INSERT INTO request_scripts (id, request_id, script_type, script) VALUES 
			(?, ?, 'pre-request', 'console.log(\"Pre-request script for post %d\");');`,
			preScriptID, requestID, post.ID)
		if err != nil {
			return fmt.Errorf("failed to insert pre-request script for request %s: %w", post.Title, err)
		}
		_, err = db.Exec(`INSERT INTO request_scripts (id, request_id, script_type, script) VALUES 
			(?, ?, 'test', 'console.log(\"Test script for post %d\");');`,
			testScriptID, requestID, post.ID)
		if err != nil {
			return fmt.Errorf("failed to insert test script for request %s: %w", post.Title, err)
		}
	}

	// Insert a sample environment.
	_, err = db.Exec(`INSERT INTO environments (id, name) VALUES
		('env-1', 'Development');`)
	if err != nil {
		return fmt.Errorf("failed to insert environment: %w", err)
	}

	// Insert realistic environment variables for JSONPlaceholder.
	_, err = db.Exec(`INSERT INTO environment_variables (id, environment_id, key, value) VALUES
		('envvar-1', 'env-1', 'BASE_URL', 'https://jsonplaceholder.typicode.com'),
		('envvar-2', 'env-1', 'HOST', 'jsonplaceholder.typicode.com'),
		('envvar-3', 'env-1', 'PORT', '443'),
		('envvar-4', 'env-1', 'PATH', '/posts');`)
	if err != nil {
		return fmt.Errorf("failed to insert environment variables: %w", err)
	}

	// Insert sample collection variables.
	_, err = db.Exec(`INSERT INTO collection_variables (id, collection_id, key, value) VALUES
		('colvar-1', 'TAPA-DEV-COL-1', 'defaultTimeout', '30000'),
		('colvar-2', 'TAPA-DEV-COL-1', 'retryAttempts', '3');`)
	if err != nil {
		return fmt.Errorf("failed to insert collection variables: %w", err)
	}

	// Insert sample sync metadata for the collection.
	_, err = db.Exec(`INSERT INTO sync_metadata (id, entity_type, entity_id, last_updated, is_dirty) VALUES
		('sync-1', 'collections', 'TAPA-DEV-COL-1', datetime('now'), 0);`)
	if err != nil {
		return fmt.Errorf("failed to insert sync metadata: %w", err)
	}

	// Insert sample keyboard shortcuts.
	_, err = db.Exec(`INSERT INTO keyboard_shortcuts (id, action, shortcut) VALUES
		('ks-1', 'Save Request', 'Ctrl+S'),
		('ks-2', 'Send Request', 'Ctrl+Enter');`)
	if err != nil {
		return fmt.Errorf("failed to insert keyboard shortcuts: %w", err)
	}

	// Insert sample user settings.
	_, err = db.Exec(`INSERT INTO user_settings (id, theme, max_history, language, font_family, font_size, working_dir) VALUES
		('singleton', 'system', 200, 'en', 'Arial', 14, '/home/user');`)
	if err != nil {
		return fmt.Errorf("failed to insert user settings: %w", err)
	}

	log.Println("Database seeded successfully.")
	return nil
}

// IsSeeded checks if the seed data exists by looking for a known record.
func IsSeeded(db *sql.DB) (bool, error) {
	var count int
	// Check for the sample collection with id "TAPA-DEV-COL-1"
	err := db.QueryRow("SELECT COUNT(1) FROM collections WHERE id = 'TAPA-DEV-COL-1'").Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check seed status: %w", err)
	}
	return count > 0, nil
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
			return fmt.Errorf("failed to flush table %s: %w", table, err)
		}
	}

	log.Println("Database flushed successfully.")
	return nil
}
