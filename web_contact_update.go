package main

import (
	"net/http"
)

func init() {
	router.PATCH("/contact", app.ContactUpdate)
}

func (app *Application) ContactUpdate(w http.ResponseWriter, r *http.Request) {
}
