package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestIndex(t *testing.T) {
	router := httprouter.New() // create router
	router.GET("/", Index)     // GET Root

	log.Fatal(http.ListenAndServe(":51234", router)) // 51234
}
