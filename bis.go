package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ConfigObject struct {
	ServiceKey string `json:"serviceKey"`
}

type Response struct {
	XMLName      xml.Name     `xml:"response"`
	ComMsgHeader ComMsgHeader `xml:"comMsgHeader"`
	MsgHeader    MsgHeader    `xml:"msgHeader"`
	MsgBody      MsgBody      `xml:"msgBody"`
}

type ComMsgHeader struct {
	XMLName    xml.Name `xml:"comMsgHeader"`
	ErrMsg     string   `xml:"errMsg"`
	ReturnCode int      `xml:"returnCode"`
}

type MsgHeader struct {
	XMLName       xml.Name `xml:"msgHeader"`
	QueryTime     string   `xml:"queryTime"`
	ResultCode    int      `xml:"resultCode"`
	ResultMessage string   `xml:"resultMessage"`
}

type MsgBody struct {
	XMLName        xml.Name         `xml:"msgBody"`
	BusStationList []BusStationList `xml:"busStationList"`
}

type BusStationList struct {
	XMLName     xml.Name `xml:"busStationList"`
	CenterYn    string   `xml:"centerYn"`
	DistrictCd  int      `xml:"districtCd"`
	MobileNo    string   `xml:"mobileNo"`
	RegionName  string   `xml:"regionName"`
	StationID   int      `xml:"stationId"`
	StationName string   `xml:"stationName"`
	X           float32  `xml:"x`
	Y           float32  `xml:"y"`
}

var config ConfigObject

// BusStopNumberToID 함수는 5자리 모바일 정류장번호를 버스 ID로 바꿔주는 기능을 합니다.
func BusStopNumberToID(number string, areaCode int) {
	file, err := ioutil.ReadFile("./configs/config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("error:", err)
	}

	URL := "http://openapi.gbis.go.kr/ws/rest/busstationservice?serviceKey=" + config.ServiceKey + "&keyword=" + number

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Response 처리
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	var busStopData Response
	xmlErr := xml.Unmarshal(data, &busStopData)
	if xmlErr != nil {
		panic(xmlErr)
	}

	fmt.Println(busStopData)
}
