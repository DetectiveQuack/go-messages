package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialiseRoutes(t *testing.T) {
	truncateTable()

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		log.Fatal(err)
	}

	body := performRequest(req)

	assert.Equal(t, "Welcome to go messages!!", string(body))
}
