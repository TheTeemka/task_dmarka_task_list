package main

import (
	"embed"

	"github.com/TheTeemka/task_dmarka_task_list/internal/config"
	"github.com/TheTeemka/task_dmarka_task_list/internal/database"
	"github.com/TheTeemka/task_dmarka_task_list/internal/repository"
	"github.com/TheTeemka/task_dmarka_task_list/internal/service"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	cfg := config.LoadConfig()

	db := database.NewPostgreSQLConnection(cfg.PSQLSource)
	taskRepo := repository.NewTaskRepo(db)
	taskService := service.NewTaskService(taskRepo)

	tagRepo := repository.NewTagsRepo(db)
	tagService := service.NewTagsService(tagRepo)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		// OnStartup:        ,
		Bind: []interface{}{
			taskService,
			tagService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
