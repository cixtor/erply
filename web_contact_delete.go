package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func init() {
	router.DELETE("/contact", app.ContactDelete)
}

// ContactDelete

//   > DELETE /contact?id=2 HTTP/1.1
//   > Host: localhost:3000
//   > Connection: close
func (app *Application) ContactDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	var id string
	var stmt *sql.Stmt

	// NOTES(cixtor): r.ParseForm doesn’t works with DELETE requests.
	if id = r.URL.Query().Get("id"); id == "" {
		fail(w, r, fmt.Errorf("missing `id` form value"))
		return
	}

	if stmt, err = app.db.Prepare("DELETE FROM contacts WHERE id=?"); err != nil {
		fail(w, r, err)
		return
	}

	if _, err = stmt.Exec(id); err != nil {
		fail(w, r, err)
		return
	}

	if _, err = contactRead(app.db, id); err == nil {
		fail(w, r, fmt.Errorf("account was not deleted"))
		return
	}

	write(w, r, Response{Ok: true})
}
