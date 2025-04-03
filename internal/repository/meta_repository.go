package repository

import (
	"context"
	"time"

	"github.com/Amir-Zouerami/TAPA/internal/models"
	"github.com/jmoiron/sqlx"
)

type MetaRepositoryInterface interface {
	IsFirstLaunch(ctx context.Context) (bool, error)
	SetFirstLaunchDone(ctx context.Context) error
	GetKeyboardShortcuts(ctx context.Context) ([]models.KeyboardShortcut, error)
	GetUserSettings(ctx context.Context) (models.UserSettings, error)
	// UpdateUserSettings(ctx context.Context, settings models.UserSettings) error
	GetLastAppState(ctx context.Context) (models.AppState, error)
	SaveLastAppState(ctx context.Context, state models.AppState) error
}

type MetaRepository struct {
	db *sqlx.DB
}

func (r *MetaRepository) IsFirstLaunch(ctx context.Context) (bool, error) {
	var firstLaunch bool
	err := r.db.GetContext(ctx, &firstLaunch, "SELECT first_launch FROM app_state WHERE id = 1")
	if err != nil {
		return false, err
	}

	return firstLaunch, nil
}

func (r *MetaRepository) SetFirstLaunchDone(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx,
		`
		UPDATE app_state 
		SET first_launch = 0, updated_at = ?
		WHERE id = 1
	`, time.Now())

	return err
}

func (r *MetaRepository) GetKeyboardShortcuts(ctx context.Context) ([]models.KeyboardShortcut, error) {
	var keyboardShortcuts []models.KeyboardShortcut
	query := `
		SELECT id, action, shortcut FROM keyboard_shortcuts
	`

	if err := r.db.SelectContext(ctx, &keyboardShortcuts, query); err != nil {
		return nil, err
	}

	return keyboardShortcuts, nil
}

func (r *MetaRepository) GetUserSettings(ctx context.Context) (models.UserSettings, error) {
	var settings models.UserSettings
	query := `SELECT id, theme, max_history, font_size FROM user_settings WHERE id = 1`

	if err := r.db.GetContext(ctx, &settings, query); err != nil {
		return models.UserSettings{}, err
	}

	return settings, nil
}

// FIXME: implement this.
// func (r *MetaRepository) UpdateUserSettings(ctx context.Context, settings models.UserSettings) error {
// 	query := `
//         INSERT INTO user_settings (id, theme, max_history, font_size)
//         VALUES (1, $1, $2, $3)
//         ON CONFLICT (id) DO UPDATE SET
//             theme = excluded.theme,
//             max_history = excluded.max_history,
//             font_size = excluded.font_size
//     `
// 	_, err := r.db.ExecContext(ctx, query, settings.Theme, settings.MaxHistory, settings.FontSize)
// 	return err
// }

func (r *MetaRepository) GetLastAppState(ctx context.Context) (models.AppState, error) {
	var state models.AppState
	err := r.db.GetContext(ctx, &state, "SELECT id, selected_environment, open_tabs, first_launch, updated_at FROM app_state WHERE id = 1")
	return state, err
}

func (r *MetaRepository) SaveLastAppState(ctx context.Context, state models.AppState) error {
	_, err := r.db.NamedExecContext(ctx, `
        UPDATE app_state 
        SET open_tabs = :open_tabs, selected_environment = :selected_environment
        WHERE id = 1
    `, state)

	return err
}

func NewMetaRepository(db *sqlx.DB) *MetaRepository {
	return &MetaRepository{db: db}
}
