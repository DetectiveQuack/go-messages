package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	var id int

	truncateTable()

	row := a.DB.QueryRow("INSERT INTO messages (message) VALUES($1) RETURNING id", "test")

	row.Scan(&id)

	route := fmt.Sprintf("/messages/%d", id)

	req, err := http.NewRequest("GET", route, nil)

	if err != nil {
		log.Fatal()
	}

	body := performRequest(req)

	assert.Equal(t, "test", string(body))
}

func TestGetMessageDb(t *testing.T) {
	truncateTable()

	req, err := http.NewRequest("GET", "/messages/10000", nil)

	if err != nil {
		log.Fatal(err)
	}

	rec := recordRequest(req)

	body, err := ioutil.ReadAll(rec.Body)

	if err != nil {
		log.Fatal()
	}

	assert.Equal(t, "sql: no rows in result set", string(body))
	assert.Equal(t, "text/plain", string(rec.Header().Get("Content-Type")))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestPostMessage(t *testing.T) {
	truncateTable()

	data := url.Values{}
	data.Set("testMessage", "")

	req, err := http.NewRequest("POST", "/messages", bytes.NewBufferString(data.Encode()))

	if err != nil {
		log.Fatal(err)
	}

	body := performRequest(req)

	assert.Equal(t, "{\"id\":1}", string(body))

	row := a.DB.QueryRow("SELECT id FROM messages WHERE id = 1")

	var id int
	row.Scan(&id)
	assert.Equal(t, id, 1)
}
