package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOnly(t *testing.T) {
	tt := []struct {
		methodType         string
		expectedStatus     string
		expectedStatusCode int
	}{
		{http.MethodGet, "StatusOK", http.StatusOK},
		{http.MethodPost, "StatusMethodNotAllowed", http.StatusMethodNotAllowed},
		{http.MethodPut, "StatusMethodNotAllowed", http.StatusMethodNotAllowed},
	}

	for _, tc := range tt {
		req, err := http.NewRequest(tc.methodType, "localhost:51234/", nil)
		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		rec := httptest.NewRecorder()

		h := GetOnly(IndexHandler)
		h(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.expectedStatusCode {
			t.Fatalf("expected %v; got %v", tc.expectedStatus, res.Status)
		}
	}
}

func TestPostOnly(t *testing.T) {
	tt := []struct {
		methodType         string
		keyword            string
		expectedStatus     string
		expectedStatusCode int
	}{
		{http.MethodPost, "아주대학교입구", "StatusOK", http.StatusOK},
		{http.MethodGet, "아주대학교입구", "StatusMethodNotAllowed", http.StatusMethodNotAllowed},
		{http.MethodPut, "아주대학교입구", "StatusMethodNotAllowed", http.StatusMethodNotAllowed},
	}

	for _, tc := range tt {
		rawBody := SearchInput{tc.keyword}

		jsonBody, err := json.Marshal(rawBody)
		if err != nil {
			t.Fatalf("could not parsed json data: %v", err)
		}
		reqBody := bytes.NewBufferString(string(jsonBody))

		req, err := http.NewRequest(tc.methodType, "localhost:51234/user/search?type=station", reqBody)
		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		rec := httptest.NewRecorder()

		h := PostOnly(SearchHandler)
		h(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.expectedStatusCode {
			t.Fatalf("expected %v; got %v", tc.expectedStatus, res.Status)
		}
	}
}
