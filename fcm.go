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
	np.Title = "승차알림"
	np.Body = fmt.Sprintf("[%v] %v번 버스가 곧 도착합니다.", GetStationNameFromStationID(routeID, stationID), GetRouteNameFromRouteID(routeID))
	np.AndroidChannelID = "stopbus_danbk_mjin1220_notification_channel_id"
	np.ClickAction = "OPEN_ACTIVITY_1"
	np.Tag = "stopbus_get_in_tag"

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
