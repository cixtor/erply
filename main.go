package main

import (
	"flag"
	"time"

	"github.com/cixtor/middleware"
)

// app is an instance of the Application, the core struct.
var app Application

// router is an instance of Middleware, a lightweight HTTP router.
var router = middleware.New()

// init functions are usually discouraged because they are called “randomly” so
// if there is more than one in the whole project, there is no guarantee that
// they will be executed in the correct order. However, when they are correctly
// used —like in this project— they allow us to provide extensibility.
func init() {
	flag.StringVar(&app.host, "host", "0.0.0.0", "Host or IP to run the server")
	flag.StringVar(&app.port, "port", "3000", "Port number to run the server")
	flag.StringVar(&app.database, "database", "database.db", "Path to the SQLite database")
	flag.StringVar(&app.migration, "migration", "migration.sql", "SQL to initialize the database")
}

func main() {
	app.Initialize()
	defer app.db.Close()

	router.Host = app.host
	router.Port = app.port

	// sane timeouts, because default Go values are crazy.
	router.IdleTimeout = 10 * time.Second
	router.ReadTimeout = 10 * time.Second
	router.WriteTimeout = 10 * time.Second
	router.ShutdownTimeout = 10 * time.Second
	router.ReadHeaderTimeout = 10 * time.Second

	router.Use(app.Auth)

	router.ListenAndServe()
}
