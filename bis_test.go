package main

import (
	"fmt"
	"net/http"
	"strings"
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

func TestGetStationIDFromStationNumber(t *testing.T) {
	tt := []struct {
		districtCd    int
		stationNumber string
		Expected      string
	}{
		{2, "04237", "203000066"},
		{2, "03126", "202000005"},
		{2, "03124", "202000039"},
		{2, "03117", "202000038"},
		{2, "03105", "202000037"},
	}
	for _, tc := range tt {
		resultData := GetStationIDFromStationNumber(tc.districtCd, tc.stationNumber)
		if resultData != tc.Expected {
			t.Fatalf("expected %v; got %v", tc.Expected, resultData)
		}
	}
}

func TestGetRouteInfo(t *testing.T) {
	tt := []struct {
		routeID    string
		httpStatus int
		resultCode int
	}{
		{"234000026", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := GetRouteInfo(tc.routeID)
		fmt.Println(resultData)
	}
}

func TestGetCurrentBusLocation(t *testing.T) {
	tt := []struct {
		routeID    string
		httpStatus int
		resultCode int
	}{
		{"234000026", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := GetCurrentBusLocation(tc.routeID)
		fmt.Println(resultData)
	}
}

func TestGetBusArrivalTime(t *testing.T) {
	tt := []struct {
		stationID  string
		httpStatus int
		resultCode int
	}{
		{"203000066", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := GetBusArrivalTime(tc.stationID)
		fmt.Println(resultData)
	}
}

func TestGetBusArrivalList(t *testing.T) {
	tt := []struct {
		stationID  string
		httpStatus int
		resultCode int
	}{
		{"203000066", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := GetBusArrivalList(tc.stationID)
		fmt.Println(resultData)
	}
}

func TestGetDataFromAPI(t *testing.T) {
	tt := []struct {
		URL               string
		expectedErrorText string
	}{
		{"http://stop-bus.tk", ""},
		{"http://stop-bus.tk/test", "Not expected http.StatusCode: 200"},
		{"http://stop-bus.tt", "no such host"},
	}
	for _, tc := range tt {
		responseBody, err := getDataFromAPI(tc.URL)

		if err != nil {
			if !strings.Contains(err.Error(), tc.expectedErrorText) {
				t.Errorf("expected %v: got %v", tc.expectedErrorText, err.Error())
			}
			continue
		}

		fmt.Println(string(responseBody))
	}
}
