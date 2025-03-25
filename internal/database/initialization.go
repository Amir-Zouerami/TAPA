package database

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "modernc.org/sqlite"

	"github.com/Amir-Zouerami/TAPA/internal/common"
	"github.com/Amir-Zouerami/TAPA/internal/config"
	"github.com/Amir-Zouerami/TAPA/internal/errors"
)

// InitializeDB creates the database file and applies its schema from db-schema.sql.
// The database is created at user config directory
func InitializeDB(schemaEmbed embed.FS) (*sqlx.DB, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, errors.Wrap(errors.ErrGetUserConfigDirectory, err)
	}

	appDir := filepath.Join(configDir, strings.ToLower(config.APP_NAME), "db")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return nil, errors.Wrap(errors.ErrCreateAppDirectory, err)
	}

	var dbPath string

	if common.IsInDevelopmentMode() {
		dbPath = filepath.Join(appDir, strings.ToLower(config.APP_NAME)+"_DEV"+".sqlite")
	} else {
		dbPath = filepath.Join(appDir, strings.ToLower(config.APP_NAME)+".sqlite")
	}

	db, err := sqlx.Open("sqlite", dbPath)
	if err != nil {
		return nil, errors.Wrap(errors.ErrOpeningDatabaseFile, err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(errors.ErrConnectingDatabase, err)
	}

	setJournalMode(db)

	if err := applySchema(schemaEmbed, db); err != nil {
		return nil, err
	}

	log.Printf("Database initialized at: %s", dbPath)
	return db, nil
}

// sets the journal mode to WAL for better concurrency
func setJournalMode(db *sqlx.DB) {
	_, err := db.Exec("PRAGMA journal_mode=WAL;")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("--> WAL MODE ENABLED")
}

// applies the database schema from the embedded db-schema.sql
func applySchema(schemaEmbed embed.FS, db *sqlx.DB) error {
	log.Println("Applying database schema...")

	rawDBSchema, err := schemaEmbed.ReadFile(config.TAPA_DB_SCHEMA_FILE_PATH)
	if err != nil {
		return errors.Wrap(errors.ErrSchemaRead, err)
	}

	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(errors.ErrSchemaCreation, err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(string(rawDBSchema))
	if err != nil {
		return errors.Wrap(errors.ErrSchemaCreation, err)
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(errors.ErrSchemaCreation, err)
	}

	log.Println("Database schema applied successfully.")
	return nil
}

func FlushAndSeedIfInDevelopmentMode(db *sqlx.DB) {
	if common.IsInDevelopmentMode() {
		err := FlushDB(db)
		if err != nil {
			log.Fatal("[TAPA_DEV_MODE] Error flushing the database:", err)
		}

		seeded, err := IsSeeded(db)
		if err != nil {
			log.Fatal("[TAPA_DEV_MODE] Error checking seed status:", err)
		}

		if !seeded {
			log.Println("[TAPA_DEV_MODE] Seeding database...")
			if err := SeedDB(db); err != nil {
				log.Fatal("Failed to seed database:", err)
			}
		} else {
			log.Println("[TAPA_DEV_MODE] Database already seeded.")
		}
	}

}
