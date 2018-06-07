package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetGoingBusList(t *testing.T) {
	retList := GetGoingBusList("228003542", "228000875")

	for _, route := range retList {
		fmt.Println(route)
	}

}

func TestSearchForStation(t *testing.T) {
	tt := []struct {
		keyword    string
		status     int
		resultCode int
	}{
		// {"운동장", http.StatusOK, 0},
		{"센터", http.StatusOK, 0},
		// {"마을", http.StatusOK, 0},
		// {"병원", http.StatusOK, 0},
	}
	for _, tc := range tt {
		// SearchForStation(tc.keyword)
		resultData := SearchForStation(tc.keyword)
		// fmt.Println(len(resultData))
		for _, rd := range resultData {
			fmt.Println(rd.StationDirect)
		}
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
		{"241005870", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := GetRouteNameFromRouteID(tc.routeID)
		fmt.Println(resultData)
	}
}

func TestGetRouteInfo(t *testing.T) {
	tt := []struct {
		routeID    string
		httpStatus int
		resultCode int
	}{
		{"241005870", http.StatusOK, 0},
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

func TestGetBusArrivalOnlyOne(t *testing.T) {
	tt := []struct {
		routeID    string
		stationID  string
		httpStatus int
		resultCode int
	}{
		{"223000100", "203000066", http.StatusOK, 0},
		{"234000026", "203000066", http.StatusOK, 0},
		{"234000024", "203000066", http.StatusOK, 0},
		{"200000053", "203000066", http.StatusOK, 0},
		{"200000110", "203000066", http.StatusOK, 0},
		{"200000112", "203000066", http.StatusOK, 0},
		{"200000144", "203000066", http.StatusOK, 0},
		{"200000064", "203000066", http.StatusOK, 0},
		{"200000146", "203000066", http.StatusOK, 0},
		{"200000070", "203000066", http.StatusOK, 0},
		{"200000119", "203000066", http.StatusOK, 0},
		{"200000208", "203000066", http.StatusOK, 0},
		{"200000211", "203000066", http.StatusOK, 0},
		{"200000231", "203000066", http.StatusOK, 0},
		{"200000185", "203000066", http.StatusOK, 0},
		{"200000236", "203000066", http.StatusOK, 0},
		{"200000272", "203000066", http.StatusOK, 0},
		{"200000196", "203000066", http.StatusOK, 0},
		{"200000197", "203000066", http.StatusOK, 0},
		{"200000199", "203000066", http.StatusOK, 0},
		{"200000201", "203000066", http.StatusOK, 0},
		{"200000205", "203000066", http.StatusOK, 0},
		{"200000320", "203000066", http.StatusOK, 0},
	}
	for _, tc := range tt {
		resultData := GetBusArrivalOnlyOne(tc.routeID, tc.stationID)
		fmt.Println(resultData)
	}
}

func TestGetBusArrivalTime(t *testing.T) {
	tt := []struct {
		stationID  string
		httpStatus int
		resultCode int
	}{
		{"115507730", http.StatusOK, 0},
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

func TestGetStationDirect(t *testing.T) {
	tt := []struct {
		stationID string
	}{
		{"202000004"},
		{"203000067"},
	}

	for _, tc := range tt {
		ret := GetStationDirect(tc.stationID)
		fmt.Println(tc.stationID,":",ret)
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
