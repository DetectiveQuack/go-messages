package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}

	a.Initialise(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	code := m.Run()

	truncateTable()

	os.Exit(code)
}

func truncateTable() {
	_, err := a.DB.Exec("TRUNCATE messages RESTART IDENTITY")

	if err != nil {
		log.Fatal()
	}
}

func recordRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()

	a.Router.ServeHTTP(rec, req)

	return rec
}

func performRequest(req *http.Request) []byte {
	rec := recordRequest(req)

	body, err := ioutil.ReadAll(rec.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
