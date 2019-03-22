package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const database = "testing.db"
const migration = "migration.sql"

func createApp(t *testing.T) Application {
	os.Remove(database)

	var app Application
	app.database = database
	app.migration = migration
	app.Initialize()

	if _, err := os.Stat(database); os.IsNotExist(err) {
		t.Fatalf("%s was not created", database)
	}

	return app
}

func TestNew(t *testing.T) {
	createApp(t)
}

func TestContactCreate(t *testing.T) {
	app := createApp(t)
	defer app.db.Close()

	var err error
	var out []byte
	var res *http.Response
	ts := httptest.NewTLSServer(http.HandlerFunc(app.ContactCreate))
	defer ts.Close()

	client := ts.Client()
	body := strings.NewReader(`firstname=John&lastname=Smith&phone=6045551234&address=350+W+Georgia+St%2C+Vancouver%2C+BC&email=john%40example.com`)
	if res, err = client.Post(ts.URL, "application/x-www-form-urlencoded", body); err != nil {
		t.Fatalf("ContactCreate; client.Post: %s", err)
		return
	}
	defer res.Body.Close()

	if out, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatalf("ContactCreate; ioutil.ReadAll: %s", err)
		return
	}

	if string(out) != `{"ok":true}`+"\n" {
		t.Fatalf("ContactCreate; failure: %s", out)
		return
	}
}

func TestContactRead(t *testing.T) {
	app := createApp(t)
	defer app.db.Close()

	var err error
	var out []byte
	var res *http.Response
	ts := httptest.NewTLSServer(http.HandlerFunc(app.ContactRead))
	defer ts.Close()

	client := ts.Client()
	if res, err = client.Get(ts.URL + "?id=1"); err != nil {
		t.Fatalf("ContactRead; client.Get: %s", err)
		return
	}
	defer res.Body.Close()

	if out, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatalf("ContactRead; ioutil.ReadAll: %s", err)
		return
	}

	if string(out) != `{"ok":true,"data":{"id":1,"firstname":"John","lastname":"Smith","phone":"6045551234","address":"350 W Georgia St, Vancouver, BC","email":"john@example.com"}}`+"\n" {
		t.Fatalf("ContactRead; failure: %s", out)
		return
	}
}
