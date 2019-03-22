package main

import (
	"net/http"
)

func init() {
	router.GET("/contact", app.ContactRead)
}

func (app *Application) ContactRead(w http.ResponseWriter, r *http.Request) {
}
