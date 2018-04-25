package StopBus

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

// Response 구조체는 GBUS API의 응답으로 오는 구조체이다.
type Response struct {
	XMLName      xml.Name     `xml:"response"`
	ComMsgHeader ComMsgHeader `xml:"comMsgHeader"`
	MsgHeader    MsgHeader    `xml:"msgHeader"`
	MsgBody      MsgBody      `xml:"msgBody"`
}

// ComMsgHeader 구조체는 GBUS API의 응답 내 공통 헤더이다.
type ComMsgHeader struct {
	XMLName    xml.Name `xml:"comMsgHeader"`
	ErrMsg     string   `xml:"errMsg"`
	ReturnCode int      `xml:"returnCode"`
}

// MsgHeader 구조체는 공통 헤더 내 메세지와 관련된 헤더이다.
type MsgHeader struct {
	XMLName       xml.Name `xml:"msgHeader"`
	QueryTime     string   `xml:"queryTime"`
	ResultCode    int      `xml:"resultCode"`
	ResultMessage string   `xml:"resultMessage"`
}

// MsgBody 구조체는  GBUS API의 응답 내 메시지 바디이다.
type MsgBody struct {
	XMLName        xml.Name     `xml:"msgBody"`
	BusStationList []BusStation `xml:"busStationList"`
}

// BusStation 구조체는 버스 정류장에 대한 정보들을 가진 구조체이다.
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

// BusStopNumberToID 함수는 5자리 모바일 정류장번호를 버스 ID로 바꿔주는 기능을 가진 함수이다.
func BusStopNumberToID(number string, areaCode int) string {
	if config.ServiceKey == "" {
		log.Println("[ERROR] config.ServiceKey not exists.")
		return ""
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

	var busStopData Response
	xmlErr := xml.Unmarshal(data, &busStopData)
	if xmlErr != nil {
		panic(xmlErr)
	}

	for i := 0; i < len(busStopData.MsgBody.BusStationList); i++ {
		if busStopData.MsgBody.BusStationList[i].DistrictCd == areaCode {
			return busStopData.MsgBody.BusStationList[i].StationID
		}
	}
	return ""
}
