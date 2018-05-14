package main

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// CommonURL is a constant that stores the common URL of the restAPI.
	CommonURL = "http://openapi.gbis.go.kr/ws/rest"
	// BusStationURLPath is a constant that stores the URL Path to the bus station.
	BusStationURLPath = "busstationservice"
	// BusRouteURLPath is a constant that stores the URL Path to the bus route.
	BusRouteURLPath = "busrouteservice"
	// BusLocationURLPath is a constant that stores the URL Path to the bus location.
	BusLocationURLPath = "buslocationservice"
)

// ComMsgHeader is a structure that specifies the data format of the common header in the APIResponseBody.
type ComMsgHeader struct {
	XMLName    xml.Name `xml:"comMsgHeader"`
	ErrMsg     string   `xml:"errMsg"`
	ReturnCode int      `xml:"returnCode"`
}

// MsgHeader is a structure that specifies the data format of the message header in the APIResponseBody.
type MsgHeader struct {
	XMLName       xml.Name `xml:"msgHeader"`
	QueryTime     string   `xml:"queryTime"`
	ResultCode    int      `xml:"resultCode"`
	ResultMessage string   `xml:"resultMessage"`
}

// StationResponseBody is a structure that specifies the data format to be responsed from the API.
type StationResponseBody struct {
	XMLName      xml.Name       `xml:"response"`
	ComMsgHeader ComMsgHeader   `xml:"comMsgHeader"`
	MsgHeader    MsgHeader      `xml:"msgHeader"`
	MsgBody      StationMsgBody `xml:"msgBody"`
}

// StationMsgBody is a structure that specifies the data format of the message body in the APIResponseBody.
type StationMsgBody struct {
	XMLName        xml.Name       `xml:"msgBody"`
	BusStationList BusStationList `xml:"busStationList"`
}

// BusStationList is an slice of BusStationes.
type BusStationList []BusStation

// BusStation is a structure that specifies the data format of the bus station in the MsgBody.
type BusStation struct {
	XMLName       xml.Name `xml:"busStationList" json:"-"`
	CenterYn      string   `xml:"centerYn" json:"-"`
	DistrictCd    int      `xml:"districtCd" json:"districtCd"`
	MobileNo      string   `xml:"mobileNo" json:"stationNumber"`
	RegionName    string   `xml:"regionName" json:"-"`
	StationID     string   `xml:"stationId" json:"stationID"`
	StationName   string   `xml:"stationName" json:"stationName"`
	X             float32  `xml:"x" json:"-"`
	Y             float32  `xml:"y" json:"-"`
	StationDirect string   `xml:"-" json:"stationDirect"`
}

// RouteResponseBody is a structure that specifies the data format to be responsed from the API.
type RouteResponseBody struct {
	XMLName      xml.Name     `xml:"response"`
	ComMsgHeader ComMsgHeader `xml:"comMsgHeader"`
	MsgHeader    MsgHeader    `xml:"msgHeader"`
	MsgBody      RouteMsgBody `xml:"msgBody"`
}

// RouteMsgBody is a structure that specifies the data format of the message body in the APIResponseBody.
type RouteMsgBody struct {
	XMLName      xml.Name     `xml:"msgBody"`
	BusRouteList BusRouteList `xml:"busRouteList"`
}

// BusRouteList is an slice of BusRoutes.
type BusRouteList []BusRoute

// BusRoute is a structure that specifies the data format of the bus route in the MsgBody.
type BusRoute struct {
	XMLName       xml.Name `xml:"busRouteList" json:"-"`
	DistrictCd    int      `xml:"districtCd" json:"districtCd"`
	RegionName    string   `xml:"regionName" json:"regionName"`
	RouteID       string   `xml:"routeId" json:"routeID"`
	RouteName     string   `xml:"routeName" json:"routeNumber"`
	RouteTypeCd   string   `xml:"routeTypeCd" json:"-"`
	RouteTypeName string   `xml:"routeTypeName" json:"routeTypeName"`
}

// RouteStationResponseBody is a structure that specifies the data format to be responsed from the API.
type RouteStationResponseBody struct {
	XMLName      xml.Name            `xml:"response"`
	ComMsgHeader ComMsgHeader        `xml:"comMsgHeader"`
	MsgHeader    MsgHeader           `xml:"msgHeader"`
	MsgBody      RouteStationMsgBody `xml:"msgBody"`
}

// RouteStationMsgBody is a structure that specifies the data format of the message body in the APIResponseBody.
type RouteStationMsgBody struct {
	XMLName             xml.Name            `xml:"msgBody"`
	BusRouteStationList BusRouteStationList `xml:"busRouteStationList"`
}

// BusRouteStationList is an slice of BusRouteStationes.
type BusRouteStationList []BusRouteStation

// BusRouteStation is a structure that specifies the data format of the bus route station in the MsgBody.
type BusRouteStation struct {
	XMLName     xml.Name `xml:"busRouteStationList" json:"-"`
	CenterYn    string   `xml:"centerYn" json:"-"`
	DistrictCd  int      `xml:"districtCd" json:"-"`
	MobileNo    string   `xml:"mobileNo" json:"stationNumber"`
	RegionName  string   `xml:"regionName" json:"-"`
	StationID   string   `xml:"stationId" json:"-"`
	StationName string   `xml:"stationName" json:"stationName"`
	X           float32  `xml:"x" json:"-"`
	Y           float32  `xml:"y" json:"-"`
	StationSeq  int      `xml:"stationSeq" json:"stationSeq"`
	TurnYn      string   `xml:"turnYn" json:"-"`
}

// RouteInfoResponseBody is a structure that specifies the data format to be responsed from the API.
type RouteInfoResponseBody struct {
	XMLName      xml.Name         `xml:"response"`
	ComMsgHeader ComMsgHeader     `xml:"comMsgHeader"`
	MsgHeader    MsgHeader        `xml:"msgHeader"`
	MsgBody      RouteInfoMsgBody `xml:"msgBody"`
}

// RouteInfoMsgBody is a structure that specifies the data format of the message body in the APIResponseBody.
type RouteInfoMsgBody struct {
	XMLName          xml.Name         `xml:"msgBody"`
	BusRouteInfoItem BusRouteInfoItem `xml:"busRouteInfoItem"`
}

// BusRouteInfoItem is a structure that specifies the data format of the bus route information in the MsgBody.
type BusRouteInfoItem struct {
	XMLName          xml.Name `xml:"busRouteInfoItem" json:"-"`
	DistrictCd       int      `xml:"districtCd" json:"districtCd"`
	DownFirstTime    string   `xml:"downFirstTime" json:"downFirstTime"`
	DownLastTime     string   `xml:"downLastTime" json:"downLastTime"`
	EndMobileNo      string   `xml:"endMobileNo" json:"endStationNumber"`
	EndStationID     string   `xml:"endStationId" json:"endStationID"`
	EndStationName   string   `xml:"endStationName" json:"endStationName"`
	RegionName       string   `xml:"regionName" json:"regionName"`
	RouteID          string   `xml:"routeId" json:"routeID"`
	RouteName        string   `xml:"routeName" json:"routeNumber"`
	RouteTypeName    string   `xml:"routeTypeName" json:"routeTypeName"`
	StartMobileNo    string   `xml:"startMobileNo" json:"startStationNumber"`
	StartStationID   string   `xml:"startStationId" json:"startStationID"`
	StartStationName string   `xml:"startStationName" json:"startStationName"`
	UpFirstTime      string   `xml:"upFirstTime" json:"upFirstTime"`
	UpLastTime       string   `xml:"upLastTime" json:"upLastTime"`
}

// LocationResponseBody is a structure that specifies the data format to be responsed from the API.
type LocationResponseBody struct {
	XMLName      xml.Name        `xml:"response"`
	ComMsgHeader ComMsgHeader    `xml:"comMsgHeader"`
	MsgHeader    MsgHeader       `xml:"msgHeader"`
	MsgBody      LocationMsgBody `xml:"msgBody"`
}

// LocationMsgBody is a structure that specifies the data format of the message body in the APIResponseBody.
type LocationMsgBody struct {
	XMLName         xml.Name        `xml:"msgBody"`
	BusLocationList BusLocationList `xml:"busLocationList"`
}

// BusLocationList is an slice of BusLocationes.
type BusLocationList []BusLocation

// BusLocation is a structure that specifies the data format of the bus location in the MsgBody.
type BusLocation struct {
	XMLName       xml.Name `xml:"busLocationList" json:"-"`
	EndBus        int      `xml:"endBus" json:"endBus"`
	LowPlate      int      `xml:"lowPlate" json:"lowPlate"`
	PlateNo       string   `xml:"plateNo" json:"plateNo"`
	RemainSeatCnt int      `xml:"remainSeatCnt" json:"remainSeatCnt"`
	StationID     string   `xml:"stationId" json:"stationId"`
	StationSeq    int      `xml:"stationSeq" json:"stationSeq"`
}

// SearchForStation is a function that searches for bus station using keywords.
func SearchForStation(keyword string) BusStationList {
	URL := CommonURL + "/" + BusStationURLPath + "?serviceKey=" + config.ServiceKey + "&keyword=" + url.PathEscape(keyword)

	responseBody, _ := getDataFromAPI(URL)

	var data StationResponseBody
	_ = xml.Unmarshal(responseBody, &data)

	return data.MsgBody.BusStationList
}

// SearchForRoute is a function that searches for bus routes using keywords.
func SearchForRoute(keyword string) BusRouteList {
	URL := CommonURL + "/" + BusRouteURLPath + "?serviceKey=" + config.ServiceKey + "&keyword=" + url.PathEscape(keyword)

	responseBody, _ := getDataFromAPI(URL)

	var data RouteResponseBody
	_ = xml.Unmarshal(responseBody, &data)

	return data.MsgBody.BusRouteList
}

// GetRouteStationList is a function that takes a list of bus line stops.
func GetRouteStationList(routeID string) BusRouteStationList {
	URL := CommonURL + "/" + BusRouteURLPath + "/station?serviceKey=" + config.ServiceKey + "&routeId=" + url.PathEscape(routeID)

	responseBody, _ := getDataFromAPI(URL)

	var data RouteStationResponseBody
	_ = xml.Unmarshal(responseBody, &data)

	return data.MsgBody.BusRouteStationList
}

// GetRouteNameFromRouteID is a function that get route name from routeID.
func GetRouteNameFromRouteID(routeID string) string {
	data := GetRouteInfo(routeID)
	return data.RouteName
}

// GetRouteInfo is a function that get route information from routeID.
func GetRouteInfo(routeID string) BusRouteInfoItem {
	URL := CommonURL + "/" + BusRouteURLPath + "/info?serviceKey=" + config.ServiceKey + "&routeId=" + url.PathEscape(routeID)

	responseBody, _ := getDataFromAPI(URL)

	var data RouteInfoResponseBody
	_ = xml.Unmarshal(responseBody, &data)

	return data.MsgBody.BusRouteInfoItem
}

// GetCurrentBusLocation is a function that takes the location of the current bus.
func GetCurrentBusLocation(routeID string) BusLocationList {
	URL := CommonURL + "/" + BusLocationURLPath + "?serviceKey=" + config.ServiceKey + "&routeId=" + url.PathEscape(routeID)

	responseBody, _ := getDataFromAPI(URL)

	var data LocationResponseBody
	_ = xml.Unmarshal(responseBody, &data)

	return data.MsgBody.BusLocationList
}

// getDataFromAPI is a function that get data from GBUS API.
func getDataFromAPI(URL string) (responseBody []byte, funcErr error) {
	response, err := http.Get(URL)
	if err != nil {
		funcErr = err
		return
	}
	if response.StatusCode != 200 {
		funcErr = errors.New("Not expected http.StatusCode: 200")
	}

	defer response.Body.Close()

	responseBody, _ = ioutil.ReadAll(response.Body)

	return
}
