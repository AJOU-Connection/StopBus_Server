package main

import (
	"fmt"

	fcm "github.com/NaySoftware/go-fcm"
)

// Alert is a function
func Alert(msg string, tokens ...string) {

	data := map[string]string{
		"msg": msg,
		"sum": "Happy Day",
	}

	ids := []string{}

	for _, token := range tokens {
		ids = append(ids, token)
	}

	xds := []string{}

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
