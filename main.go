package main

import (
	"embed"
	"log"

	"github.com/Amir-Zouerami/TAPA/internal/config"
	"github.com/Amir-Zouerami/TAPA/internal/database"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist build/tapa.png
var assets embed.FS

//go:embed internal/database/db-schema.sql
var dbSchema embed.FS

func main() {
	app := config.NewApp()

	db, err := database.InitializeDB(dbSchema)
	if err != nil {
		log.Fatal("Database initialization failed \n", err)
	}

	defer db.Close()

	appConfig, err := config.GetAppConfig(assets, app)
	if err != nil {
		log.Fatal("Could not bootstrap the application \n", err)
	}

	err = wails.Run(&options.App{
		Title:            appConfig.Title,
		Width:            appConfig.Width,
		Height:           appConfig.Height,
		WindowStartState: appConfig.WindowStartState,
		AssetServer:      appConfig.AssetServer,
		Linux:            appConfig.Linux,
		Mac:              appConfig.Mac,
		OnStartup:        appConfig.OnStartup,
		Bind:             appConfig.Bind,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
