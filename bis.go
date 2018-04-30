package StopBus

import (
	"encoding/xml"
	"fmt"
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

// APIResponseBody is a structure that specifies the data format to be responsed from the API.
type APIResponseBody struct {
	XMLName      xml.Name     `xml:"response"`
	ComMsgHeader ComMsgHeader `xml:"comMsgHeader"`
	MsgHeader    MsgHeader    `xml:"msgHeader"`
	MsgBody      MsgBody      `xml:"msgBody"`
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

// MsgBody is a structure that specifies the data format of the message body in the APIResponseBody.
type MsgBody struct {
	XMLName        xml.Name     `xml:"msgBody"`
	BusStationList []BusStation `xml:"busStationList"`
}

// BusStationList is an slice of BusStationes.
type BusStationList []BusStation

// BusStation is a structure that specifies the data format of the bus stastion in the MsgBody.
type BusStation struct {
	XMLName     xml.Name `xml:"busStationList"`
	CenterYn    string   `xml:"centerYn"`
	DistrictCd  int      `xml:"districtCd"`
	MobileNo    string   `xml:"mobileNo"`
	RegionName  string   `xml:"regionName"`
	StationID   string   `xml:"stationId"`
	StationName string   `xml:"stationName"`
	X           float32  `xml:"x"`
	Y           float32  `xml:"y"`
}

// SearchForStation is a function that searches for bus station using keywords.
func SearchForStation(keyword string) {
	URL := CommonURL + "/" + BusStationURLPath + "?serviceKey=" + config.ServiceKey + "&keyword=" + url.PathEscape(keyword)

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var data APIResponseBody
	xmlErr := xml.Unmarshal(responseBody, &data)
	if xmlErr != nil {
		panic(xmlErr)
	}

	fmt.Println(data)
}

// SearchForRoute is a function that searches for bus routes using keywords.
func SearchForRoute(keyword string) {
	URL := CommonURL + "/" + BusRouteURLPath + "?serviceKey=" + config.ServiceKey + "&keyword=" + url.PathEscape(keyword)

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var data APIResponseBody
	xmlErr := xml.Unmarshal(responseBody, &data)
	if xmlErr != nil {
		panic(xmlErr)
	}

	fmt.Println(data)
}
