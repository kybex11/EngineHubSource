package main

import (
	"embed"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/hugolgst/rich-go/client"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type ActivityConfig struct {
	State   string `json:"state"`
	Details string `json:"details"`
}

var assets embed.FS

func main() {
	app := NewApp()

	erro := client.Login("831160231148781598")
	if erro != nil {
		panic(erro)
	}

	file, err := ioutil.ReadFile("info.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var config ActivityConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	erro = client.SetActivity(client.Activity{
		State:   config.State,
		Details: config.Details,
	})

	if erro != nil {
		panic(erro)
	}

	err = wails.Run(&options.App{
		Title:     "Engine Hub",
		Width:     1024,
		Height:    768,
		MinHeight: 768,
		MinWidth:  1024,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Println("Error:", err.Error())
	}
}
