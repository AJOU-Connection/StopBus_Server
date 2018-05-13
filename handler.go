package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// JSONBody is a structure that specifies the JSON format to put in the body of the response packet.
type JSONBody struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

// The Header is a structure that specifies the Header part of the JSONBody.
type Header struct {
	Result       bool   `json:"result"`
	ErrorCode    int    `json:"errorCode"`
	ErrorContent string `json:"errorContent"`
}

// BusStationInput is a structure that specifies the format of POST Response Body in DriverRegisterHandler.
type BusStationInput struct {
	BusNumber           string      `json:"busNumber"`
	BusRouteStationList interface{} `json:"stationList"`
}

// DriverInput is a structure that specifies the format of POST Request Body in DriverRegisterHandler.
type DriverInput struct {
	PlateNo string `json:"plateNo"`
	RouteID string `json:"routeID"`
}

// SearchInput is a structure that specifies the format of POST Request Body in SearchHandler.
type SearchInput struct {
	Keyword string `json:"keyword"`
}

// Handler is a function that handles the entire routing in the server.
func Handler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", GetOnly(IndexHandler))
	r.HandleFunc("/driver/register", PostOnly(DriverRegisterHandler))
	r.HandleFunc("/user/search", PostOnly(SearchHandler))

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

// IndexHandler is a function that handles routing for index access.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "StopBus")
}

// DriverRegisterHandler is a function that handles the routing for bus driver registration.
func DriverRegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var di DriverInput
	_ = json.Unmarshal(body, &di)

	BISData := GetRouteStationList(di.RouteID)

	busNumber := GetRouteNameFromRouteID(di.RouteID)

	jsonBody := JSONBody{
		Header{true, 0, ""},
		BusStationInput{busNumber, BISData},
	}
	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

// SearchHandler is a function that handles routing for bus route or stop searching.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var si SearchInput
	_ = json.Unmarshal(body, &si)

	header := Header{true, 0, ""}
	var data interface{}

	searchType := r.FormValue("type")
	if searchType == "route" {
		data = SearchForRoute(si.Keyword)
	} else if searchType == "station" {
		data = SearchForStation(si.Keyword)
	} else {
		header.Result = false
		header.ErrorCode = 1
		header.ErrorContent = "Invalid Search Type: " + searchType
	}

	jsonBody := JSONBody{
		header,
		data,
	}

	var jsonValue []byte

	if data == nil {
		jsonValue, _ = json.Marshal(struct {
			Header Header `json:"header"`
		}{jsonBody.Header})
	} else {
		jsonValue, _ = json.Marshal(jsonBody)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}
