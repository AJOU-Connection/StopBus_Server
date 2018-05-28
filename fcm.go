package main

import (
	"fmt"

	fcm "github.com/NaySoftware/go-fcm"
)

// GetInAlert is a function
func GetInAlert(routeID string, stationID string, tokens ...string) {
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

	data["Title"] = "승차알림"
	data["Message"] = fmt.Sprintf("[%v] %v번 버스가 곧 도착합니다.", GetStationNameFromStationID(routeID, stationID), GetRouteNameFromRouteID(routeID))

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
