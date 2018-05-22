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

type OnlyStationNumberInput struct {
	DistrictCd    int    `json:"districtCd"`
	StationNumber string `json:"stationNumber"`
}
type User struct {
	Token string `json:"token"`
	UUID  string `json:"UUID"`
}

type Reserv struct {
	UserToken string `json:userToken`
	RouteID   string `json:routeID`
	StationID string `json:stationID`
}

// Handler is a function that handles the entire routing in the server.
func Handler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", GetOnly(IndexHandler))
	r.HandleFunc("/driver/register", PostOnly(DriverRegisterHandler))
	r.HandleFunc("/driver/gap", PostOnly(GapHandler))
	r.HandleFunc("/user/register", PostOnly(UserRegisterHandler))
	r.HandleFunc("/user/search", PostOnly(SearchHandler))
	r.HandleFunc("/user/routeInfo", PostOnly(RouteInfoHandler))
	r.HandleFunc("/user/busLocationList", PostOnly(BusLocationListHandler))
	r.HandleFunc("/user/busStationList", PostOnly(BusStationListHandler))
	r.HandleFunc("/user/busArrival", PostOnly(BusArrivalHandler))
	r.HandleFunc("/user/reserv/getIn", PostOnly(ReservGetInHandler))
	r.HandleFunc("/user/reserv/getOut", PostOnly(ReservGetInHandler))

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

// GapHandler is a function that handles the routing for gap time between buses
func GapHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var gi GapInput
	_ = json.Unmarshal(body, &gi)

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

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var user User
	_ = json.Unmarshal(body, &user)

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

// RouteInfoHandler is a function that handles routing for bus route information
func RouteInfoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var orii OnlyRouteIDInput
	_ = json.Unmarshal(body, &orii)

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

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var orii OnlyRouteIDInput
	_ = json.Unmarshal(body, &orii)

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

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var orii OnlyRouteIDInput
	_ = json.Unmarshal(body, &orii)

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

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var osni OnlyStationNumberInput
	_ = json.Unmarshal(body, &osni)

	header := Header{true, 0, ""}
	var data interface{}

	stationID := GetStationIDFromStationNumber(osni.DistrictCd, osni.StationNumber)
	data = GetBusArrivalTime(stationID)

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

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	var reserv Reserv
	_ = json.Unmarshal(body, &reserv)

	jsonBody := struct {
		Header Header `json:"header"`
	}{
		Header{true, 0, ""},
	}

	ret := addGetIn(reserv)
	if ret != nil {
		jsonBody.Header.Result = false
		jsonBody.Header.ErrorCode = 1
		jsonBody.Header.ErrorContent = "Failed to reserve get in"
	}

	jsonValue, _ := json.Marshal(jsonBody)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonValue))
}

// func ReservGetOutHandler(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()

// 	body, _ := ioutil.ReadAll(r.Body)
// 	fmt.Println(string(body))

// 	var reserv Reserv
// 	_ = json.Unmarshal(body, &reserv)

// 	jsonBody := struct {
// 		Header Header `json:"header"`
// 	}{
// 		Header{true, 0, ""},
// 	}

// 	ret := addGetOut(reserv)
// 	if ret != nil {
// 		jsonBody.Header.Result = false
// 		jsonBody.Header.ErrorCode = 1
// 		jsonBody.Header.ErrorContent = "Failed to reserve getIn"
// 	}

// 	jsonValue, _ := json.Marshal(jsonBody)

// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprintln(w, string(jsonValue))
// }
