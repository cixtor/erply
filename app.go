package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Application is the base struct for the whole web API service.
//
// All the API endpoints are attached to this struct as methods, this allows us
// to reference additional resources like the database, cache layer (if there
// is one), and other middlewares.
type Application struct {
	// db is the connection with the SQLite database.
	db *sql.DB

	// host is the name or IP address for the web server.
	host string

	// port is the port number where the web server is running.
	port string

	// database is the full path to the SQLite database file.
	database string
}

func (app *Application) Init(database string) {
	var err error
	var db *sql.DB

	if db, err = sql.Open("sqlite3", database); err != nil {
		fmt.Println("sql.Open", err)
		os.Exit(1)
	}

	if _, err = db.Exec(dbtables); err != nil {
		fmt.Println("db.Init", err)
		os.Exit(1)
	}

	app.db = db
}
