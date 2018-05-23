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
	}

	ids := []string{}

	for _, token := range tokens {
		ids = append(ids, token)
	}

	xds := []string{}

	np := fcm.NotificationPayload{}
	np.Title = "testTitle"
	np.Body = "곧 버스가 도착합니다."
	np.AndroidChannelID = "stopbus_danbk_mjin1220_notification_channel_id"
	np.ClickAction = "OPEN_ACTIVITY_1"

	c := fcm.NewFcmClient(config.ServerKey)
	c.SetNotificationPayload(&np)
	c.NewFcmRegIdsMsg(ids, data)
	c.AppendDevices(xds)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
