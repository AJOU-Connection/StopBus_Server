package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type JSONBody struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

type Header struct {
	Result       bool   `json:"result"`
	ErrorCode    int    `json:"errorCode"`
	ErrorContent string `json:"errorContent"`
}

type BusStationInfo struct {
	BusNumber           string                 `json:"busNumber"`
	BusRouteStationList ResBusRouteStationList `json:"stationList"`
}
type ResBusRouteStationList []ResBusRouteStaion
type ResBusRouteStaion struct {
	MobileNo    string `json:"stationNumber"`
	StationName string `json:"stationName"`
	StationSeq  int    `json:"stationSeq"`
}

type DriverInfo struct {
	PlateNo string `json:"plateNo"`
	RouteID string `json:"routeID"`
}

func Handler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", GetOnly(IndexHandler))
	r.HandleFunc("/driver/register", PostOnly(DriverRegisterHandler))
	r.HandleFunc("/user/search", PostOnly(SearchHandler))

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "StopBus")
}

func DriverRegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var di DriverInfo
	_ = json.Unmarshal(body, &di)

	BISData := GetRouteStationList(di.RouteID)
	resDataList := ResBusRouteStationList{}
	for _, data := range BISData {
		resDataList = append(resDataList, ResBusRouteStaion{strings.TrimSpace(data.MobileNo), data.StationName, data.StationSeq})
	}
	busNumber := GetRouteNameFromRouteID(di.RouteID)

	jsonBody := JSONBody{
		Header{true, 0, ""},
		BusStationInfo{busNumber, resDataList},
	}
	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	header := Header{true, 0, ""}
	data := ""

	searchType := r.FormValue("type")
	if searchType == "route" {
		data = "route"
	} else if searchType == "station" {
		data = "station"
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

	if data == "" {
		jsonValue, _ = json.Marshal(struct {
			Header Header `json:"header"`
		}{jsonBody.Header})
	} else {
		jsonValue, _ = json.Marshal(jsonBody)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}