package main

import (
	"log"
	"net/http"
)

// main is the main function.
func main() {
	err := http.ListenAndServe(":51234", Handler())
	if err != nil {
		log.Fatal(err)
	}
}
