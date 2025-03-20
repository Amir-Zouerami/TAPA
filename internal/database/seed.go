package database

import (
	"database/sql"
	"fmt"
	"log"
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

	if err := seedCollection(db); err != nil {
		return err
	}

	if err := seedFolders(db); err != nil {
		return err
	}

	if err := seedRequests(db); err != nil {
		return err
	}

	if err := seedEnvironment(db); err != nil {
		return err
	}

	if err := seedCollectionVariables(db); err != nil {
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

func seedCollection(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO collections (id, name, description) VALUES
		('TAPA-DEV-COL-1', 'Sample Collection', 'A collection to store API requests and related data');`)

	if err != nil {
		return fmt.Errorf("failed to insert collection: %w", err)
	}

	return nil
}

func seedFolders(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO folders (id, collection_id, name, position) VALUES
		('TAPA-DEV-FOLDER-1', 'TAPA-DEV-COL-1', 'General Requests', 1);`)

	if err != nil {
		return fmt.Errorf("failed to insert folder: %w", err)
	}

	return nil
}

// fakes a request to JSONPlaceholder
func seedRequests(db *sql.DB) error {
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

func seedRequestAndDependencies(db *sql.DB, post PlaceholderPost) error {
	requestID := fmt.Sprintf("req-%d", post.ID)
	_, err := db.Exec(`INSERT INTO requests 
		(id, collection_id, folder_id, name, method, url, body, timeout, allow_redirects, ssl_verification, remove_referer_on_redirect, encode_url)
		VALUES (?, 'TAPA-DEV-COL-1', 'TAPA-DEV-FOLDER-1', ?, 'GET', ?, '', 30000, 1, 1, 0, 1);`,
		requestID, post.Title, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", post.ID))

	if err != nil {
		return fmt.Errorf("failed to insert request %s: %w", post.Title, err)
	}

	if err := seedRequestHeaders(db, requestID, post); err != nil {
		return err
	}

	if err := seedRequestQueryParams(db, requestID, post); err != nil {
		return err
	}

	if post.ID == 1 {
		if err := seedRequestAuth(db, requestID, post); err != nil {
			return err
		}
	}

	if err := seedRequestHistory(db, requestID, post); err != nil {
		return err
	}

	if err := seedRequestScripts(db, requestID, post); err != nil {
		return err
	}

	return nil
}

func seedRequestHeaders(db *sql.DB, requestID string, post PlaceholderPost) error {
	headerID := fmt.Sprintf("hdr-%d", post.ID)
	_, err := db.Exec(`INSERT INTO request_headers (id, request_id, key, value) VALUES 
		(?, ?, 'Content-Type', 'application/json');`, headerID, requestID)

	if err != nil {
		return fmt.Errorf("failed to insert header for request %s: %w", post.Title, err)
	}

	return nil
}

func seedRequestQueryParams(db *sql.DB, requestID string, post PlaceholderPost) error {
	queryParamID := fmt.Sprintf("qry-%d", post.ID)
	_, err := db.Exec(`INSERT INTO request_query_params (id, request_id, key, value) VALUES 
		(?, ?, 'userId', '1');`, queryParamID, requestID)

	if err != nil {
		return fmt.Errorf("failed to insert query parameter for request %s: %w", post.Title, err)
	}

	return nil
}

func seedRequestAuth(db *sql.DB, requestID string, post PlaceholderPost) error {
	_, err := db.Exec(`INSERT INTO request_auth (id, request_id, auth_type, token, username, password) VALUES 
		('auth-1', ?, 'Bearer', 'sample-token-123', '', '');`, requestID)

	if err != nil {
		return fmt.Errorf("failed to insert auth for request %s: %w", post.Title, err)
	}

	return nil
}

func seedRequestHistory(db *sql.DB, requestID string, post PlaceholderPost) error {
	historyID := fmt.Sprintf("hist-%d", post.ID)
	now := time.Now().Format(time.DateTime)

	_, err := db.Exec(`INSERT INTO request_history 
		(id, request_id, timestamp, history_type, method, url, headers, query_params, auth, body, status_code, response, response_headers)
		VALUES (?, ?, ?, 'auto', 'GET', ?, '{"Content-Type":"application/json"}', '{"userId":"1"}', '', '', 200, '{"id":1}', '{"Content-Length":"123"}');`,
		historyID, requestID, now, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", post.ID))

	if err != nil {
		return fmt.Errorf("failed to insert history for request %s: %w", post.Title, err)
	}

	return nil
}

func seedRequestScripts(db *sql.DB, requestID string, post PlaceholderPost) error {
	preScriptID := fmt.Sprintf("script-pre-%d", post.ID)
	testScriptID := fmt.Sprintf("script-test-%d", post.ID)

	_, err := db.Exec(`INSERT INTO request_scripts (id, request_id, script_type, script) VALUES 
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
	return nil
}

func seedEnvironment(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO environments (id, name) VALUES
		('env-1', 'Development');`)
	if err != nil {
		return fmt.Errorf("failed to insert environment: %w", err)
	}

	_, err = db.Exec(`INSERT INTO environment_variables (id, environment_id, key, value) VALUES
		('envvar-1', 'env-1', 'BASE_URL', 'https://jsonplaceholder.typicode.com'),
		('envvar-2', 'env-1', 'HOST', 'jsonplaceholder.typicode.com'),
		('envvar-3', 'env-1', 'PORT', '443'),
		('envvar-4', 'env-1', 'PATH', '/posts');`)
	if err != nil {
		return fmt.Errorf("failed to insert environment variables: %w", err)
	}
	return nil
}

func seedCollectionVariables(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO collection_variables (id, collection_id, key, value) VALUES
		('colvar-1', 'TAPA-DEV-COL-1', 'defaultTimeout', '30000'),
		('colvar-2', 'TAPA-DEV-COL-1', 'retryAttempts', '3');`)

	if err != nil {
		return fmt.Errorf("failed to insert collection variables: %w", err)
	}

	return nil
}

func seedSyncMetadata(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO sync_metadata (id, entity_type, entity_id, last_updated, is_dirty) VALUES
		('sync-1', 'collections', 'TAPA-DEV-COL-1', datetime('now'), 0);`)

	if err != nil {
		return fmt.Errorf("failed to insert sync metadata: %w", err)
	}

	return nil
}

func seedKeyboardShortcuts(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO keyboard_shortcuts (id, action, shortcut) VALUES
		('ks-1', 'Save Request', 'Ctrl+S'),
		('ks-2', 'Send Request', 'Ctrl+Enter');`)

	if err != nil {
		return fmt.Errorf("failed to insert keyboard shortcuts: %w", err)
	}

	return nil
}

func seedUserSettings(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO user_settings (id, theme, max_history, language, font_family, font_size, working_dir) VALUES
		('singleton', 'system', 200, 'en', 'Arial', 14, '/home/user');`)

	if err != nil {
		return fmt.Errorf("failed to insert user settings: %w", err)
	}

	return nil
}

// IsSeeded checks if the seed data exists by looking for a known record.
func IsSeeded(db *sql.DB) (bool, error) {
	var count int
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
