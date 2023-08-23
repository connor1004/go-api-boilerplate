package main

import (
	"os"

	"github.com/connor1004/go-api-boilerplate/app"
)

func main() {
	appMain := app.NewApp()
	appMain.InitializeDB()
	defer appMain.DB.Close()

	args := os.Args
	if len(args) > 1 && args[1] == "import" {
		appMain.ImportUsers()
		return
	}

	appMain.InitializeRoutes()
	appMain.Run()
}
