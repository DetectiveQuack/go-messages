// Go app to store and retrieve messages

package main

import (
	"log"
	"net/http"
	"os"
)

// Initialise db and serve app
func main() {
	a := App{}

	a.Initialise(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	log.Fatal(http.ListenAndServe(":3001", a.Router))
}
