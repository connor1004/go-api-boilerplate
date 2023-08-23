package app

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"

	"github.com/connor1004/go-api-boilerplate/config"
	"github.com/connor1004/go-api-boilerplate/controllers"
	"github.com/connor1004/go-api-boilerplate/utils"
	_ "github.com/go-sql-driver/mysql" // mysql-driver
)

// Handler - route handler function type
type Handler func(*utils.Context)

// Route - a route struct with Pattern and a handler
type Route struct {
	Pattern *regexp.Regexp
	Method  string
	Handler Handler
}

// App -
type App struct {
	DB           *sql.DB
	Routes       []Route
	DefaultRoute Handler
}

// NewApp - create a new App
func NewApp() *App {
	app := &App{
		DefaultRoute: func(ctx *utils.Context) {
			ctx.Respond(http.StatusNotFound, map[string]interface{}{"message": "Not found"})
		},
	}

	return app
}

// InitializeDB - initialize the App
func (a *App) InitializeDB() {
	connectionString := config.DBConfig["user"] + ":" + config.DBConfig["password"] +
		"@/" + config.DBConfig["dbName"]

	db, err := sql.Open(config.DBConfig["driver"], connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
}

// InitializeRoutes - initialize routes for the app
func (a *App) InitializeRoutes() {
	userController := controllers.NewUserController(a.DB)
	a.Handle(`^/api/users$`, "POST", userController.AddUser)

	a.Handle(`^/api/users/([\w\._-]+)$`, "GET", userController.GetUserByID)

	a.Handle(`^/api/search-users$`, "GET", userController.SearchUsers)
}

// Run - run the app
func (a *App) Run() {
	err := http.ListenAndServe(":8080", a)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

// Handle - process a route pattern and add it to app.Routes
func (a *App) Handle(pattern string, method string, handler Handler) {
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Method: method, Handler: handler}

	a.Routes = append(a.Routes, route)
}

// ServeHTTP - route Handler function
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.Context{Request: r, ResponseWriter: w}

	for _, rt := range a.Routes {
		if matches := rt.Pattern.FindStringSubmatch(r.URL.Path); len(matches) > 0 && r.Method == rt.Method {
			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			rt.Handler(ctx)
			return
		}
	}

	a.DefaultRoute(ctx)
}
