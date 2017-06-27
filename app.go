package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"

	"log"

	"net/http"

	"github.com/gorilla/mux"
)

// App struct holds app configuration
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialise database
func (a *App) Initialise(user, password, db string) {
	con := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, db)

	var err error
	a.DB, err = sql.Open("postgres", con)

	if err != nil {
		log.Fatal(err)
	}

	err = a.DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.InitialiseRoutes()
}

// InitialiseRoutes sets up messages routes
func (a *App) InitialiseRoutes() {
	a.InitialiseMessages()

	a.Router.HandleFunc("/", getRoot)
}

// getRoot sends message on root
func getRoot(w http.ResponseWriter, r *http.Request) {
	SendPlainText(w, "Welcome to go messages!!", http.StatusOK)
}
