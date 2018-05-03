package main

import (
	"fmt"
	"net/http"
)

func Handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", IndexHandler)
	return r
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "StopBus")
}
