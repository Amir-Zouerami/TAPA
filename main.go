package main

import (
	"embed"
	"log"

	"github.com/Amir-Zouerami/TAPA/internal/config"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist build/tapa.png
var assets embed.FS

func main() {
	app := config.NewApp()
	config, err := config.GetAppConfig(assets, app)

	if err != nil {
		log.Fatal("Could not bootstrap the application \n", err)
	}

	err = wails.Run(&options.App{
		Title:            config.Title,
		Width:            config.Width,
		Height:           config.Height,
		WindowStartState: config.WindowStartState,
		AssetServer:      config.AssetServer,
		Linux:            config.Linux,
		Mac:              config.Mac,
		OnStartup:        config.OnStartup,
		Bind:             config.Bind,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
