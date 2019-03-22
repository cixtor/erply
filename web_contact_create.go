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
//   > Host: localhost:3000
//   > Content-Length: 115
//   >
//   > firstname=John&lastname=Smith&phone=6045551234&address=350+W+Georgia+St%2C+Vancouver%2C+BC&email=john%40example.com
func (app *Application) ContactCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var stmt *sql.Stmt

	if err = r.ParseForm(); err != nil {
		fmt.Println("r.ParseForm", err)
		fail(w, r, fmt.Errorf("r.ParseForm %s", err))
		return
	}

	firstname := r.Form.Get("firstname")
	lastname := r.Form.Get("lastname")
	phone := r.Form.Get("phone")
	address := r.Form.Get("address")
	email := r.Form.Get("email")

	if stmt, err = app.db.Prepare(`INSERT INTO contacts (firstname, lastname, phone, address, email) VALUES(?, ?, ?, ?, ?)`); err != nil {
		fmt.Println("db.Prepare", err)
		fail(w, r, fmt.Errorf("db.Prepare %s", err))
		return
	}

	if _, err = stmt.Exec(firstname, lastname, phone, address, email); err != nil {
		fmt.Println("stmt.Exec", err)
		fail(w, r, fmt.Errorf("stmt.Exec %s", err))
		return
	}

	write(w, r, Response{Ok: true})
}
