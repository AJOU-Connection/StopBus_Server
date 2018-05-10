package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSearchForStation(t *testing.T) {
	tt := []struct {
		keyword    string
		status     int
		resultCode int
	}{
		{"아주대학교입구", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := SearchForStation(tc.keyword)
		fmt.Println(resultData)
	}
}

func TestSearchForRoute(t *testing.T) {
	tt := []struct {
		keyword    string
		status     int
		resultCode int
	}{
		{"3007", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := SearchForRoute(tc.keyword)
		fmt.Println(resultData)
	}
}

func TestGetRouteStationList(t *testing.T) {
	tt := []struct {
		routeID    string
		httpStatus int
		resultCode int
	}{
		{"234000026", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := GetRouteStationList(tc.routeID)
		fmt.Println(resultData)
	}
}

func TestGetRouteNameFromRouteID(t *testing.T) {
	tt := []struct {
		routeID    string
		httpStatus int
		resultCode int
	}{
		{"234000026", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := GetRouteNameFromRouteID(tc.routeID)
		fmt.Println(resultData)
	}
}

func TestGetDataFromAPI(t *testing.T) {
	tt := []struct {
		URL               string
		expectedErrorText string
	}{
		{"http://stop-bus.tk", ""},
		{"http://stop-bus.tk/test", "Not expected http.StatusCode: 200."},
		{"http://stop-bus.tt", "Get http://stop-bus.tt: dial tcp: lookup stop-bus.tt: no such host"},
	}
	for _, tc := range tt {
		responseBody, err := getDataFromAPI(tc.URL)

		if err != nil {
			if err.Error() != tc.expectedErrorText {
				t.Errorf("expected %v: got %v", tc.expectedErrorText, err.Error())
			}
			continue
		}

		fmt.Println(string(responseBody))
	}
}
