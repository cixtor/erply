package main

import (
	"net/http"
)

func init() {
	router.DELETE("/contact", app.ContactDelete)
}

func (app *Application) ContactDelete(w http.ResponseWriter, r *http.Request) {
}
