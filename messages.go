package main

import (
	"net/http"

	"database/sql"

	"strconv"

	"io/ioutil"

	"github.com/gorilla/mux"
)

type message struct {
	ID      int `json:"id"`
	message string
	sRouter *mux.Router
}

// InitialiseMessages initilialises message routes, prefix router
func (a *App) InitialiseMessages() {
	m := message{}

	m.sRouter = a.Router.PathPrefix("/messages").Subrouter()

	m.sRouter.HandleFunc("/{id:[0-9]+}", a.getMessage).Methods("GET")
	m.sRouter.HandleFunc("", a.postMessage).Methods("POST")
}

// Perform validation checks and get the message from the db
func (a *App) getMessage(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	m := message{}

	err := m.getMessage(a.DB, id)

	if err != nil {
		SendPlainText(w, err.Error(), http.StatusBadRequest)
		return
	}

	SendPlainText(w, m.message, http.StatusOK)
}

// Get the message from the database using the id passed in
func (m *message) getMessage(db *sql.DB, id int) (err error) {

	row := db.QueryRow("SELECT message FROM messages WHERE ID=$1", id)

	return row.Scan(&m.message)
}

// postMessage inserts a message to the db and returns an id
func (a *App) postMessage(w http.ResponseWriter, r *http.Request) {
	m := message{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		SendJSON(w, message{message: "Error reading post message"}, http.StatusBadRequest)
		return
	}

	err = m.postMessage(a.DB, string(body))

	if err != nil {
		SendJSON(w, message{message: "DB post error"}, http.StatusBadRequest)
		return
	}

	SendJSON(w, m, http.StatusOK)
}

// postMessage inserts message into database
func (m *message) postMessage(db *sql.DB, text string) (err error) {
	row := db.QueryRow("INSERT INTO messages (message) VALUES($1) RETURNING id", text)

	return row.Scan(&m.ID)
}
