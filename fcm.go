package main

import (
	"fmt"
	"log"

	fcm "github.com/NaySoftware/go-fcm"
)

// GetInAlert is a function
func GetInAlert(routeID string, stationID string) {
	tokens, err := getGetInUserTokens(routeID, stationID)
	if err != nil {
		fmt.Println(err)
	}

	title := "승차알림"
	message := fmt.Sprintf("[%v] %v번 버스가 곧 도착합니다.", GetStationNameFromStationID(routeID, stationID), GetRouteNameFromRouteID(routeID))
	fcmAlert(routeID, stationID, title, message, tokens)
}

func GetOutAlert(routeID string, stationID string, plateNo string) {
	tokens, err := getGetOutUserTokens(routeID, stationID, plateNo)
	if err != nil {
		fmt.Println(err)
	}

	title := "하차알림"
	message := fmt.Sprintf("%v번 버스가 곧 %v에 도착합니다.", GetRouteNameFromRouteID(routeID), GetStationNameFromStationID(routeID, stationID))
	fcmAlert(routeID, stationID, title, message, tokens)
}

func GetInAlertUsingUUID(reserv Reserv) {
	title := "승차알림"
	message := fmt.Sprintf("[%v] %v번 버스가 곧 도착합니다.", GetStationNameFromStationID(reserv.RouteID, reserv.StationID), GetRouteNameFromRouteID(reserv.RouteID))

	token, err := getTokenFromUUID(reserv.UUID)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	fcmAlert(reserv.RouteID, reserv.StationID, title, message, []string{token})
}

func fcmAlert(routeID string, stationID string, title string, message string, tokens []string) {
	data := map[string]string{
		"routeID":   routeID,
		"stationID": stationID,
		"Title":     "",
		"Message":   "",
	}

	ids := []string{}

	for _, token := range tokens {
		ids = append(ids, token)
	}

	xds := []string{}

	data["Title"] = title
	data["Message"] = message

	c := fcm.NewFcmClient(config.ServerKey)
	c.NewFcmRegIdsMsg(ids, data)
	c.AppendDevices(xds)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
