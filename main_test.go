package main

import (
	"os"
	"testing"

	"github.com/connor1004/go-api-boilerplate/app"
)

var testApp *app.App

func TestMain(m *testing.M) {
	testApp = app.NewApp()
	testApp.InitializeDB()
	defer testApp.DB.Close()

	testApp.InitializeRoutes()

	code := m.Run()

	os.Exit(code)
}
