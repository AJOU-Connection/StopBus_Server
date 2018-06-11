package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

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

type StopInput GapInput
type GapInput struct {
	RouteID   string `json:"routeID"`
	StationID string `json:"stationID"`
}

// SearchInput is a structure that specifies the format of POST Request Body in SearchHandler.
type SearchInput struct {
	Keyword string `json:"keyword"`
}

// OnlyRouteIDInput is a structure that specifies the format of the POST Request Body with only the RouteID.
type OnlyRouteIDInput struct {
	RouteID string `json:"routeID"`
}

type OnlyStationIDInput struct {
	StationID string `json:"stationID"`
}
type User struct {
	Token string `json:"token"`
	UUID  string `json:"UUID"`
}

type Reserv struct {
	UUID      string `json:"UUID"`
	RouteID   string `json:"routeID"`
	StationID string `json:"stationID"`
	PlateNo   string `json:"plateNo"`
}

type GetInfo struct {
	IsGetIn  bool `json:"isGetIn"`
	IsGetOff bool `json:"isGetOff"`
}

type IsGoInfo struct {
	SourceStationID string `json:"sourceStationID"`
	DestiStationID  string `json:"destiStationID"`
}

type StarInfo struct {
	RouteIDList []string `json:"routeIDList"`
}

type StationInfo struct {
	StationID     string `json:"stationID"`
	StationNumber string `json:"stationNumber"`
}

// Handler is a function that handles the entire routing in the server.
func Handler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", GetOnly(IndexHandler))

	r.HandleFunc("/driver/register", PostOnly(DriverRegisterHandler))
	r.HandleFunc("/driver/gap", PostOnly(DriverGapHandler))
	r.HandleFunc("/driver/stop", PostOnly(DriverStopHandler))

	r.HandleFunc("/user/register", PostOnly(UserRegisterHandler))
	r.HandleFunc("/user/routeInfo", PostOnly(RouteInfoHandler))
	r.HandleFunc("/user/starInfo", PostOnly(StarInfoHandler))
	r.HandleFunc("/user/search", PostOnly(SearchHandler))
	r.HandleFunc("/user/stationName", PostOnly(StationNameHandler))
	r.HandleFunc("/user/stationDirect", PostOnly(StationDirectHandler))
	r.HandleFunc("/user/busLocationList", PostOnly(BusLocationListHandler))
	r.HandleFunc("/user/busStationList", PostOnly(BusStationListHandler))
	r.HandleFunc("/user/busArrival", PostOnly(BusArrivalHandler))
	r.HandleFunc("/user/reserv/getIn", PostOnly(ReservGetInHandler))
	r.HandleFunc("/user/reserv/getOut", PostOnly(ReservGetOutHandler))

	r.HandleFunc("/user/reserv/panel", PostOnly(ReservPanelHandler))
	r.HandleFunc("/user/isgo", PostOnly(IsGoHandler))

	loggedRouter := handlers.LoggingHandler(io.Writer(GetLogFile()), r)
	return loggedRouter
}

// IndexHandler is a function that handles routing for index access.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "StopBus")
}

// DriverRegisterHandler is a function that handles the routing for bus driver registration.
func DriverRegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var di DriverInput
	decodeJSON(r.Body, &di)

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

// DriverGapHandler is a function that handles the routing for gap time between buses
func DriverGapHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var gi GapInput
	decodeJSON(r.Body, &gi)

	BISData := GetBusArrivalTime(gi.StationID)
	var gapData interface{}

	for _, data := range BISData {
		if data.RouteID == gi.RouteID {
			gapData = struct {
				LocationNo1  int    `json:"locationNo1"`
				LocationNo2  int    `json:"locationNo2"`
				PlateNo1     string `json:"plateNo1"`
				PlateNo2     string `json:"plateNo2"`
				PredictTime1 int    `json:"predictTime1"`
				PredictTime2 int    `json:"predictTime2"`
			}{
				data.LocationNo1,
				data.LocationNo2,
				data.PlateNo1,
				data.PlateNo2,
				data.PredictTime1,
				data.PredictTime2,
			}
		}
	}

	jsonBody := JSONBody{
		Header{true, 0, ""},
		gapData,
	}
	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func DriverStopHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var si StopInput
	decodeJSON(r.Body, &si)

	jsonBody := JSONBody{
		Header{true, 0, ""},
		GetInfo{},
	}

	getInfo, err := getGetCount(si.RouteID, si.StationID)
	if err != nil {
		jsonBody.Header.Result = false
		jsonBody.Header.ErrorCode = 1
		jsonBody.Header.ErrorContent = "Failed to select DriverStop count"
	} else {
		jsonBody.Body = getInfo
	}

	var jsonValue []byte
	if err != nil {
		onlyHeaderJSON := struct {
			Header Header `json:"header"`
		}{
			jsonBody.Header,
		}

		jsonValue, _ = json.Marshal(onlyHeaderJSON)
	} else {
		jsonValue, _ = json.Marshal(jsonBody)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user User
	decodeJSON(r.Body, &user)

	jsonBody := struct {
		Header Header `json:"header"`
	}{
		Header{true, 0, ""},
	}

	ret := addUserToken(user)
	if ret != nil {
		jsonBody.Header.Result = false
		jsonBody.Header.ErrorCode = 1
		jsonBody.Header.ErrorContent = "Failed to add user"
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func StarInfoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var starInfo StarInfo
	decodeJSON(r.Body, &starInfo)

	reqLength := len(starInfo.RouteIDList)

	header := Header{true, 0, ""}
	data := make([]BusRouteInfoItem, reqLength, reqLength)

	var wg sync.WaitGroup
	for index, routeID := range starInfo.RouteIDList {
		wg.Add(1)

		go func(routeID string, index int, wg *sync.WaitGroup) {
			defer wg.Done()
			info := GetRouteInfo(routeID)
			data[index] = info
		}(routeID, index, &wg)
	}
	wg.Wait()

	jsonBody := JSONBody{
		header,
		data,
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

// SearchHandler is a function that handles routing for bus route or stop searching.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var si SearchInput
	decodeJSON(r.Body, &si)

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

// RouteInfoHandler is a function that handles routing for bus route information
func RouteInfoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var orii OnlyRouteIDInput
	decodeJSON(r.Body, &orii)

	header := Header{true, 0, ""}
	var data interface{}

	data = GetRouteInfo(orii.RouteID)

	jsonBody := JSONBody{
		header,
		data,
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

// BusStationListHandler is a function that handles routing for bus station list
func BusStationListHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var orii OnlyRouteIDInput
	decodeJSON(r.Body, &orii)

	header := Header{true, 0, ""}
	var data interface{}

	data = GetRouteStationList(orii.RouteID)

	jsonBody := JSONBody{
		header,
		data,
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

// BusLocationListHandler is a function that handles routing for bus location list
func BusLocationListHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var orii OnlyRouteIDInput
	decodeJSON(r.Body, &orii)

	header := Header{true, 0, ""}
	var data interface{}

	data = GetCurrentBusLocation(orii.RouteID)

	jsonBody := JSONBody{
		header,
		data,
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

// BusArrivalHandler is a function that handles routing for bus arrival time
func BusArrivalHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var osii OnlyStationIDInput
	decodeJSON(r.Body, &osii)

	header := Header{true, 0, ""}
	var data interface{}

	data = GetBusArrivalTime(osii.StationID)

	jsonBody := JSONBody{
		header,
		data,
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func ReservGetInHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reserv Reserv
	decodeJSON(r.Body, &reserv)

	jsonBody := struct {
		Header Header `json:"header"`
	}{
		Header{true, 0, ""},
	}

	busList := GetBusArrivalList(reserv.StationID)
	if !isInBusList(reserv.RouteID, busList) {
		jsonBody.Header.Result = false
		jsonBody.Header.ErrorCode = 2
		jsonBody.Header.ErrorContent = "Invalid bus routeID and stationID"
	} else {
		data := GetBusArrivalOnlyOne(reserv.RouteID, reserv.StationID)

		if data.PlateNo1 == "" {
			goto END
		}
		isFirstInput, err := addDriverStop(StopInput{reserv.RouteID, reserv.StationID}, GetIn)
		if err != nil {
			jsonBody.Header.Result = false
			jsonBody.Header.ErrorCode = 3
			jsonBody.Header.ErrorContent = "Failed to reserve get in"
			goto END
		} else if isFirstInput {
			go isBusPassed(reserv.RouteID, reserv.StationID)
		}

		if data.PredictTime1 <= 2 {
			go GetInAlertUsingUUID(reserv)
		} else {
			ret := addGetIn(reserv)
			if ret != nil {
				jsonBody.Header.Result = false
				jsonBody.Header.ErrorCode = 1
				jsonBody.Header.ErrorContent = "Failed to reserve get in"
				goto END
			}
			if isFirstInput {
				go TargetObserver(reserv.RouteID, reserv.StationID)
			}
		}
	}

END:
	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func ReservGetOutHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reserv Reserv
	decodeJSON(r.Body, &reserv)

	jsonBody := struct {
		Header Header `json:"header"`
	}{
		Header{true, 0, ""},
	}

	busList := GetBusArrivalList(reserv.StationID)
	if !isInBusList(reserv.RouteID, busList) {
		jsonBody.Header.Result = false
		jsonBody.Header.ErrorCode = 3
		jsonBody.Header.ErrorContent = "Invalid bus routeID and stationID"
	} else {
		locationData := GetCurrentBusLocation(reserv.RouteID)

		if !isInCurrentBusList(reserv, locationData) {
			jsonBody.Header.Result = false
			jsonBody.Header.ErrorCode = 2
			jsonBody.Header.ErrorContent = "There is no bus in service."
			goto END
		}

		ret := addGetOut(reserv)
		if ret != nil {
			jsonBody.Header.Result = false
			jsonBody.Header.ErrorCode = 1
			jsonBody.Header.ErrorContent = "Failed to reserve get out"
			goto END
		}

		go isTargetBusPassed(reserv.RouteID, reserv.StationID, reserv.PlateNo)
	}

END:
	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func IsGoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var isgo IsGoInfo
	decodeJSON(r.Body, &isgo)

	header := Header{true, 0, ""}
	var data interface{}

	data = GetGoingBusList(isgo.SourceStationID, isgo.DestiStationID)

	jsonBody := JSONBody{
		header,
		data,
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func ReservPanelHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reserv Reserv
	decodeJSON(r.Body, &reserv)

	jsonBody := struct {
		Header Header `json:"header"`
	}{
		Header{true, 0, ""},
	}

	isFirstInput, err := addDriverStop(StopInput{reserv.RouteID, reserv.StationID}, GetIn)
	if err != nil {
		jsonBody.Header.Result = false
		jsonBody.Header.ErrorCode = 1
		jsonBody.Header.ErrorContent = "Failed to reserve get in at panel"
		log.Printf("%v\n", err)
	} else if isFirstInput {
		go isBusPassed(reserv.RouteID, reserv.StationID)
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func StationNameHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var stationInfo StationInfo
	decodeJSON(r.Body, &stationInfo)

	jsonBody := JSONBody{
		Header{true, 0, ""},
		struct {
			StationName   string `json:"stationName"`
			StationDirect string `json:"stationDirect"`
		}{
			GetStationName(stationInfo.StationNumber, stationInfo.StationID),
			GetStationDirect(stationInfo.StationID),
		},
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func StationDirectHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var stationIDInput OnlyStationIDInput
	decodeJSON(r.Body, &stationIDInput)

	data := GetStationDirect(stationIDInput.StationID)

	jsonBody := JSONBody{
		Header{true, 0, ""},
		struct {
			StationDirect string `json:"stationDirect"`
		}{
			data,
		},
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

func decodeJSON(r io.Reader, subject interface{}) {
	if err := json.NewDecoder(r).Decode(subject); err != nil {
		log.Fatalln(err)
	}
}

func isInBusList(routeID string, busList BusRouteList) bool {
	ret := false
	for _, bus := range busList {
		if bus.RouteID == routeID {
			ret = true
			break
		}
	}
	return ret
}

func isInCurrentBusList(reserv Reserv, busLocationList BusLocationList) bool {
	ret := false
	currentStaOrder := 0

	for _, bus := range busLocationList {
		if bus.PlateNo[len(bus.PlateNo)-4:] == reserv.PlateNo {
			currentStaOrder = bus.StationSeq
			ret = true
			break
		}
	}

	if (reserv.RouteID != "201320974") && (GetBusArrivalOnlyOne(reserv.RouteID, reserv.StationID).StaOrder < currentStaOrder) {
		ret = false
	}

	return ret
}
