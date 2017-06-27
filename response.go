package main

import (
	"encoding/json"
	"net/http"
)

const (
	errorMsg = `Something went wrong!!!`
)

// SendPlainText sends plain text message
func SendPlainText(w http.ResponseWriter, text string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(text))
}

// SendJSON sends JSON response
func SendJSON(w http.ResponseWriter, payload interface{}, code int) {
	res, err := json.Marshal(payload)

	if err != nil {
		res = []byte(errorMsg)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(code)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
