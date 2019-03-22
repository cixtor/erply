package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func init() {
	router.PATCH("/contact", app.ContactUpdate)
}

// ContactUpdate updates an existing contact in the database.
//
//   > PATCH /contact HTTP/1.1
//   > Content-Type: application/x-www-form-urlencoded; charset=utf-8
//   > Authorization: Basic Zm9vQGV4YW1wbGUuY29tOnBhc3N3b3Jk
//   > Host: localhost:3000
//   > Connection: close
//   >
//   > id=1&firstname=Yorman&lastname=Arias&phone=6045551234&address=350+W+Georgia+St%2C+Vancouver%2C+BC&email=john%40example.com
func (app *Application) ContactUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var id string
	var stmt *sql.Stmt

	if err = r.ParseForm(); err != nil {
		fail(w, r, fmt.Errorf("r.ParseForm %s", err))
		return
	}

	if id = r.Form.Get("id"); id == "" {
		fail(w, r, fmt.Errorf("missing `id` form value"))
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

	if stmt, err = app.db.Prepare("UPDATE contacts SET firstname=?, lastname=?, phone=?, address=?, email=? WHERE id=?"); err != nil {
		fail(w, r, err)
		return
	}

	if _, err = stmt.Exec(
		c.Firstname,
		c.Lastname,
		c.Phone,
		c.Address,
		c.Email,
		id,
	); err != nil {
		fail(w, r, err)
		return
	}

	if c, err = contactRead(app.db, id); err != nil {
		fail(w, r, err)
		return
	}

	write(w, r, Response{Ok: true, Data: c})
}
