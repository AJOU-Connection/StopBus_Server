package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", GetOnly(IndexHandler))
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "StopBus")
}
