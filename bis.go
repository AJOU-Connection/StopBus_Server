package main

import (
	"encoding/xml"
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
)

// StationResponseBody is a structure that specifies the data format to be responsed from the API.
type StationResponseBody struct {
	XMLName      xml.Name       `xml:"response"`
	ComMsgHeader ComMsgHeader   `xml:"comMsgHeader"`
	MsgHeader    MsgHeader      `xml:"msgHeader"`
	MsgBody      StationMsgBody `xml:"msgBody"`
}

// RouteResponseBody is a structure that specifies the data format to be responsed from the API.
type RouteResponseBody struct {
	XMLName      xml.Name     `xml:"response"`
	ComMsgHeader ComMsgHeader `xml:"comMsgHeader"`
	MsgHeader    MsgHeader    `xml:"msgHeader"`
	MsgBody      RouteMsgBody `xml:"msgBody"`
}

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

// StationMsgBody is a structure that specifies the data format of the message body in the APIResponseBody.
type StationMsgBody struct {
	XMLName        xml.Name       `xml:"msgBody"`
	BusStationList BusStationList `xml:"busStationList"`
}

// BusStationList is an slice of BusStationes.
type BusStationList []BusStation

// BusStation is a structure that specifies the data format of the bus stastion in the MsgBody.
type BusStation struct {
	XMLName       xml.Name `xml:"busStationList" json:"-"`
	CenterYn      string   `xml:"centerYn" json:"-"`
	DistrictCd    int      `xml:"districtCd" json:"districtCd"`
	MobileNo      string   `xml:"mobileNo" json:"stationNumber"`
	RegionName    string   `xml:"regionName" json:"-"`
	StationID     string   `xml:"stationId" json:"-"`
	StationName   string   `xml:"stationName" json:"stationName"`
	X             float32  `xml:"x" json:"-"`
	Y             float32  `xml:"y" json:"-"`
	StationDirect string   `xml:"-" json:"stationDirect"`
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
	RegionName    string   `xml:"regionName" json:"-"`
	RouteID       string   `xml:"routeId" json:"-"`
	RouteName     string   `xml:"routeName" json:"routeNumber"`
	RouteTypeCd   string   `xml:"routeTypeCd" json:"-"`
	RouteTypeName string   `xml:"routeTypeName" json:"routeTypeName"`
}

type RouteStationResponseBody struct {
	XMLName      xml.Name            `xml:"response"`
	ComMsgHeader ComMsgHeader        `xml:"comMsgHeader"`
	MsgHeader    MsgHeader           `xml:"msgHeader"`
	MsgBody      RouteStationMsgBody `xml:"msgBody"`
}

type RouteStationMsgBody struct {
	XMLName             xml.Name            `xml:"msgBody"`
	BusRouteStationList BusRouteStationList `xml:"busRouteStationList"`
}
type BusRouteStationList []BusRouteStation
type BusRouteStation struct {
	XMLName     xml.Name `xml:"busRouteStationList"`
	CenterYn    string   `xml:"centerYn"`
	DistrictCd  int      `xml:"districtCd"`
	MobileNo    string   `xml:"mobileNo"`
	RegionName  string   `xml:"regionName"`
	StationID   string   `xml:"stationId"`
	StationName string   `xml:"stationName"`
	X           float32  `xml:"x"`
	Y           float32  `xml:"y"`
	StationSeq  int      `xml:"stationSeq"`
	TurnYn      string   `xml:"turnYn"`
}

// SearchForStation is a function that searches for bus station using keywords.
func SearchForStation(keyword string) BusStationList {
	URL := CommonURL + "/" + BusStationURLPath + "?serviceKey=" + config.ServiceKey + "&keyword=" + url.PathEscape(keyword)

	responseBody := getDataFromAPI(URL)

	var data StationResponseBody
	_ = xml.Unmarshal(responseBody, &data)

	return data.MsgBody.BusStationList
}

// SearchForRoute is a function that searches for bus routes using keywords.
func SearchForRoute(keyword string) BusRouteList {
	URL := CommonURL + "/" + BusRouteURLPath + "?serviceKey=" + config.ServiceKey + "&keyword=" + url.PathEscape(keyword)

	responseBody := getDataFromAPI(URL)

	var data RouteResponseBody
	_ = xml.Unmarshal(responseBody, &data)

	return data.MsgBody.BusRouteList
}

func GetRouteStationList(routeID string) BusRouteStationList {
	URL := CommonURL + "/" + BusRouteURLPath + "/station?serviceKey=" + config.ServiceKey + "&routeId=" + url.PathEscape(routeID)

	responseBody := getDataFromAPI(URL)

	var data RouteStationResponseBody
	xmlErr := xml.Unmarshal(responseBody, &data)
	if xmlErr != nil {
		panic(xmlErr)
	}

	return data.MsgBody.BusRouteStationList
}

func GetRouteNameFromRouteID(routeID string) string {
	URL := CommonURL + "/" + BusRouteURLPath + "/info?serviceKey=" + config.ServiceKey + "&routeId=" + url.PathEscape(routeID)

	responseBody := getDataFromAPI(URL)

	var data struct {
		XMLName      xml.Name     `xml:"response"`
		ComMsgHeader ComMsgHeader `xml:"comMsgHeader"`
		MsgHeader    MsgHeader    `xml:"msgHeader"`
		MsgBody      struct {
			XMLName          xml.Name `xml:"msgBody"`
			BusRouteInfoItem struct {
				XMLName   xml.Name `xml:"busRouteInfoItem"`
				RouteName string   `xml:"routeName"`
			} `xml:busRouteInfoItem`
		} `xml:msgBody`
	}
	xmlErr := xml.Unmarshal(responseBody, &data)
	if xmlErr != nil {
		panic(xmlErr)
	}

	return data.MsgBody.BusRouteInfoItem.RouteName
}

func getDataFromAPI(URL string) []byte {
	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return responseBody
}
