package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
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

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
}

func TestDriverGapHandler(t *testing.T) {
	tt := []struct {
		gapInput       GapInput
		httpStatusCode int
	}{
		{GapInput{"234000026", "201000348"}, http.StatusOK},
		{GapInput{"234000026", "201000126"}, http.StatusOK},
		{GapInput{"234000026", "201000125"}, http.StatusOK},
		{GapInput{"234000026", "202000015"}, http.StatusOK},
		{GapInput{"234000026", "203000066"}, http.StatusOK},
	}

	for _, tc := range tt {
		jsonBody, err := json.Marshal(tc.gapInput)
		if err != nil {
			t.Fatalf("could not parsed json data: %v", err)
		}
		reqBody := bytes.NewBufferString(string(jsonBody))

		req, err := http.NewRequest("POST", "localhost:51234/driver/gap", reqBody)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		DriverGapHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.httpStatusCode {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
	}
}

func TestUserRegisterHandler(t *testing.T) {
	for i := 0; i < 10; i++ {
		rawBody := User{"testToken_" + strconv.Itoa(time.Now().Nanosecond()), "testUUID_" + strconv.Itoa(time.Now().Nanosecond())}

		jsonBody, err := json.Marshal(rawBody)
		if err != nil {
			t.Fatalf("could not parsed json data: %v", err)
		}
		reqBody := bytes.NewBufferString(string(jsonBody))

		req, err := http.NewRequest("POST", "localhost:51234/user/register", reqBody)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		UserRegisterHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
	}
}

func TestSearchHandler(t *testing.T) {
	tt := []struct {
		reqType        string
		keyword        string
		httpStatusCode int
	}{
		{"station", "아주대학교입구", http.StatusOK},
		{"station", "고척시장", http.StatusOK},
		{"route", "720-2", http.StatusOK},
		{"stat", "", http.StatusOK},
	}

	for _, tc := range tt {
		rawBody := SearchInput{tc.keyword}

		jsonBody, err := json.Marshal(rawBody)
		if err != nil {
			t.Fatalf("could not parsed json data: %v", err)
		}
		reqBody := bytes.NewBufferString(string(jsonBody))

		req, err := http.NewRequest("POST", "localhost:51234/user/search?type="+tc.reqType, reqBody)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		SearchHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.httpStatusCode {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
	}
}

func TestRouteInfoHandler(t *testing.T) {
	tt := []struct {
		routeID        string
		httpStatusCode int
	}{
		{"234000026", http.StatusOK},
		{"232000092", http.StatusOK},
		{"210000039", http.StatusOK},
		{"234000021", http.StatusOK},
		{"222000002", http.StatusOK},
	}

	for _, tc := range tt {
		rawBody := OnlyRouteIDInput{tc.routeID}

		jsonBody, err := json.Marshal(rawBody)
		if err != nil {
			t.Fatalf("could not parsed json data: %v", err)
		}
		reqBody := bytes.NewBufferString(string(jsonBody))

		req, err := http.NewRequest("POST", "localhost:51234/user/routeInfo", reqBody)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		RouteInfoHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.httpStatusCode {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
	}
}

func TestBusLocationListHandler(t *testing.T) {
	tt := []struct {
		routeID        string
		httpStatusCode int
	}{
		{"234000026", http.StatusOK},
		{"232000092", http.StatusOK},
		{"210000039", http.StatusOK},
		{"234000021", http.StatusOK},
		{"222000002", http.StatusOK},
	}

	for _, tc := range tt {
		rawBody := OnlyRouteIDInput{tc.routeID}

		jsonBody, err := json.Marshal(rawBody)
		if err != nil {
			t.Fatalf("could not parsed json data: %v", err)
		}
		reqBody := bytes.NewBufferString(string(jsonBody))

		req, err := http.NewRequest("POST", "localhost:51234/user/busLocationList", reqBody)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		BusLocationListHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.httpStatusCode {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
	}
}

func TestBusStationListHandler(t *testing.T) {
	tt := []struct {
		routeID        string
		httpStatusCode int
	}{
		{"234000026", http.StatusOK},
		{"232000092", http.StatusOK},
		{"210000039", http.StatusOK},
		{"234000021", http.StatusOK},
		{"222000002", http.StatusOK},
	}

	for _, tc := range tt {
		rawBody := OnlyRouteIDInput{tc.routeID}

		jsonBody, err := json.Marshal(rawBody)
		if err != nil {
			t.Fatalf("could not parsed json data: %v", err)
		}
		reqBody := bytes.NewBufferString(string(jsonBody))

		req, err := http.NewRequest("POST", "localhost:51234/user/busStationList", reqBody)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		BusStationListHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.httpStatusCode {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
	}
}

func TestBusArrivalHandler(t *testing.T) {
	tt := []struct {
		stationID      string
		httpStatusCode int
	}{
		{"203000066", http.StatusOK},
	}

	for _, tc := range tt {
		rawBody := OnlyStationIDInput{tc.stationID}

		jsonBody, err := json.Marshal(rawBody)
		if err != nil {
			t.Fatalf("could not parsed json data: %v", err)
		}
		reqBody := bytes.NewBufferString(string(jsonBody))

		req, err := http.NewRequest("POST", "localhost:51234/user/busArrival", reqBody)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		BusArrivalHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.httpStatusCode {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
	}
}
