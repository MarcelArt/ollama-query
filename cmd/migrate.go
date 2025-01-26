package cmd

import (
	"os"

	"github.com/MarcelArt/ollama-query/config"
	"github.com/MarcelArt/ollama-query/database"
)

func Migrate(arg string) {
	switch arg {
	case "up":
		config.SetupENV()
		database.ConnectDB()
		database.MigrateDB()
		database.SeedDB()
	case "down":
		config.SetupENV()
		database.ConnectDB()
		database.DropDB()
	default:
		println("Unknown command")
		os.Exit(0)
	}
}
