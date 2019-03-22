package main

import (
	"net/http"
)

func init() {
	router.POST("/contact", app.ContactCreate)
}

func (app *Application) ContactCreate(w http.ResponseWriter, r *http.Request) {
}
