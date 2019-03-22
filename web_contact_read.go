package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func init() {
	router.GET("/contact", app.ContactRead)
}

// CreateRead returns a single contact from the database, if available.
//
//   > GET /contact?id=1 HTTP/1.1
//   > Host: localhost:3000
//   > Connection: close
func (app *Application) ContactRead(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	c, err := contactRead(app.db, id)

	if err != nil {
		fail(w, r, err)
		return
	}

	write(w, r, Response{Ok: true, Data: c})
}

func contactRead(db *sql.DB, id string) (Contact, error) {
	var c Contact
	var err error
	var stmt *sql.Stmt

	if id == "" {
		return Contact{}, fmt.Errorf("missing `id` query parameter")
	}

	if stmt, err = db.Prepare("SELECT * FROM contacts WHERE id = ?"); err != nil {
		return Contact{}, err
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
		return Contact{}, err
	}

	return c, nil
}
