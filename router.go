package StopBus

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ResponseJSON 구조체는 서버의 응답에 넣어 보낼 구조체이다
type ResponseJSON struct {
	Header Header `json:"header"`
}

// Header 구조체는 ResponseJSON의 Header 구조체이다.
type Header struct {
	Result        bool   `json:"result"`
	ErrorContents string `json:"errorContents"`
}

// Index 함수는 index경로로 접속했을 때의 라우트 함수이다.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json") // header json setting
	w.WriteHeader(http.StatusOK)                       // create http header

	responseJSON := ResponseJSON{Header{true, ""}} // create json variable

	json.NewEncoder(w).Encode(responseJSON) // send json with given struct
}
