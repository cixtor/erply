package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func init() {
	router.POST("/contact", app.ContactCreate)
}

// ContactCreate inserts a new contact into the database.
//
//   > POST /contact HTTP/1.1
//   > Content-Type: application/x-www-form-urlencoded; charset=utf-8
//   > Authorization: Basic Zm9vQGV4YW1wbGUuY29tOnBhc3N3b3Jk
//   > Host: localhost:3000
//   > Content-Length: 115
//   >
//   > firstname=John&lastname=Smith&phone=6045551234&address=350+W+Georgia+St%2C+Vancouver%2C+BC&email=john%40example.com
func (app *Application) ContactCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var stmt *sql.Stmt

	if err = r.ParseForm(); err != nil {
		fail(w, r, fmt.Errorf("r.ParseForm %s", err))
		return
	}

	c := Contact{
		Firstname: r.Form.Get("firstname"),
		Lastname:  r.Form.Get("lastname"),
		Phone:     r.Form.Get("phone"),
		Address:   r.Form.Get("address"),
		Email:     r.Form.Get("email"),
	}

	if err = c.Valid(); err != nil {
		fail(w, r, err)
		return
	}

	if stmt, err = app.db.Prepare(`INSERT INTO contacts (firstname, lastname, phone, address, email) VALUES(?, ?, ?, ?, ?)`); err != nil {
		fail(w, r, fmt.Errorf("db.Prepare %s", err))
		return
	}

	if _, err = stmt.Exec(
		c.Firstname,
		c.Lastname,
		c.Phone,
		c.Address,
		c.Email,
	); err != nil {
		fail(w, r, fmt.Errorf("stmt.Exec %s", err))
		return
	}

	write(w, r, Response{Ok: true})
}
