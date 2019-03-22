package main

import (
	"os"
	"testing"
)

const database = "testing.db"

func TestInit(t *testing.T) {
	var app Application
	app.Init(database)
	app.db.Close()

	if _, err := os.Stat(database); os.IsNotExist(err) {
		t.Fatalf("%s should have been created, but was not", database)
	}

	os.Remove(database)
}
