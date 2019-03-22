package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	// migration is an SQL file to initialize the database.
	migration string
}

// Response represents the data to respond to XHR requests.
type Response struct {
	Ok    bool        `json:"ok"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// Initialize creates the database tables and inserts initial data.
func (app *Application) Initialize() {
	var err error
	var db *sql.DB

	if db, err = sql.Open("sqlite3", app.database); err != nil {
		router.Logger.Println("sql.Open", err)
		os.Exit(1)
	}

	app.db = db

	// load SQL migration file, if available.
	if out, err2 := ioutil.ReadFile(app.migration); err2 == nil {
		if _, err = db.Exec(string(out)); err != nil {
			router.Logger.Println("db.Init", err)
			os.Exit(1)
		}
	}
}

// write writes a JSON encoded object with a successful message.
func write(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(v); err != nil {
		router.Logger.Println("json.Encode", err)
		return
	}
}

// fail writes a JSON encoded object with an error message.
func fail(w http.ResponseWriter, r *http.Request, err error) {
	write(w, r, Response{Error: err.Error()})
}
