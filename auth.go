package main

import (
	"database/sql"
	"net/http"
)

// Auth handles the authentication of all the API endpoints.
func (app *Application) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		var id int
		var ok bool
		var err error
		var stmt *sql.Stmt
		var user string
		var pass string

		if user, pass, ok = r.BasicAuth(); !ok {
			http.Error(w, http.StatusText(401), http.StatusUnauthorized)
			return
		}

		// NOTES(cixtor): as explained in the README, password is not really a
		// user password, this is a plain API Authentication Key, we don’t need
		// to salt nor hash this text as “security” is out of the scope of this
		// project.
		if stmt, err = app.db.Prepare("SELECT id FROM keys WHERE username = ? AND password = ?"); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			router.Logger.Println("AuthError", err)
			return
		}

		defer stmt.Close()

		if err = stmt.QueryRow(user, pass).Scan(&id); err != nil {
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}

		w.Header().Del("WWW-Authenticate")

		next.ServeHTTP(w, r)
	})
}
