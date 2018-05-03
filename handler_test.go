package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/", srv.URL))
	if err != nil{
		t.Fatalf("could not send GET request: %v", err)
	}
	
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	data := string(bytes.TrimSpace(body))
	if data != "StopBus" {
		t.Fatalf("expected StopBus; got %v", data)
	}

}

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:51234/", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()

	IndexHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	data := string(bytes.TrimSpace(body))
	if data != "StopBus" {
		t.Fatalf("expected StopBus; got %v", data)
	}
}
