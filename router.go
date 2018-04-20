package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ResponseJSON is response format struct
type ResponseJSON struct {
	Header Header `json:"header"`
}

// Header is response header format struct
type Header struct {
	Result        bool   `json:"result"`
	ErrorContents string `json:"errorContents"`
}

// Index is a route function that operate as index page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json") // header json setting
	w.WriteHeader(http.StatusOK)                       // create http header

	responseJSON := ResponseJSON{Header{true, ""}} // create json variable

	json.NewEncoder(w).Encode(responseJSON) // send json with given struct
}
