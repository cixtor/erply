package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Contact represents a single record in the database.
type Contact struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}

func init() {
	router.GET("/contact", app.ContactRead)
}

// CreateRead returns a single contact from the database, if available.
//
//   > GET /contact?id=1 HTTP/1.1
//   > Host: localhost:3000
//   > Connection: close
func (app *Application) ContactRead(w http.ResponseWriter, r *http.Request) {
	var c Contact
	var id string
	var err error
	var stmt *sql.Stmt

	if id = r.URL.Query().Get("id"); id == "" {
		fail(w, r, fmt.Errorf("missing `id` query parameter"))
		return
	}

	if stmt, err = app.db.Prepare("SELECT * FROM contacts WHERE id = ?"); err != nil {
		fail(w, r, err)
		return
	}

	defer stmt.Close()

	if err = stmt.QueryRow(id).Scan(
		&c.ID,
		&c.Firstname,
		&c.Lastname,
		&c.Phone,
		&c.Address,
		&c.Email,
	); err != nil {
		fail(w, r, err)
		return
	}

	write(w, r, Response{Ok: true, Data: c})
}
