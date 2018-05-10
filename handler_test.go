package main

import (
	"bytes"
	"encoding/json"
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
	if err != nil {
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

func TestDriverRegisterHandler(t *testing.T) {
	rawBody := DriverInput{"경기00가1234", "234000026"}

	jsonBody, err := json.Marshal(rawBody)
	if err != nil {
		t.Fatalf("could not parsed json data: %v", err)
	}
	reqBody := bytes.NewBufferString(string(jsonBody))

	req, err := http.NewRequest("POST", "localhost:51234/driver/register", reqBody)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	DriverRegisterHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", res.Status)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	fmt.Println(string(resBody))
}
