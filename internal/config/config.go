package config

import (
	"context"
	"io/fs"

	"github.com/Amir-Zouerami/TAPA/internal/common"
	"github.com/Amir-Zouerami/TAPA/internal/errors"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

// appConfig holds application settings.
type appConfig struct {
	Title            string
	Width            int
	Height           int
	Frameless        bool
	WindowStartState options.WindowStartState
	AssetServer      *assetserver.Options
	Linux            *linux.Options
	Mac              *mac.Options
	Bind             []any
	OnStartup        func(ctx context.Context)
}

// GetAppConfig creates application configuration with default settings.
func GetAppConfig(assets fs.FS, app *App) (*appConfig, error) {
	icon, err := common.ReadEmbeddedFile(assets, TAPA_ICON_NAME)

	if err != nil {
		return nil, errors.Wrap(errors.ErrAppConfigGeneration, err)
	}

	return &appConfig{
		Title:            APP_NAME,
		Width:            APP_WIDTH,
		Height:           APP_HEIGHT,
		WindowStartState: options.Maximised,
		OnStartup:        app.startup,
		Linux: &linux.Options{
			Icon: icon,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{Title: APP_NAME, Message: MACOS_ABOUT_MESSAGE},
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []any{
			app,
		},
	}, nil
}
