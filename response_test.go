package main

import (
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestSendJSON(t *testing.T) {
	j := make(chan int)
	rec := httptest.NewRecorder()

	SendJSON(rec, j, http.StatusOK)

	body, err := ioutil.ReadAll(rec.Body)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "Something went wrong!!!", string(body))
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
