package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ResponseJson struct {
	result string
}

// Index 함수는 '/'로 접속했을 때의 서버 동작을 포함하고 있다.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rj := ResponseJson{"response"}

	fmt.Fprintln(w, rj)
}
