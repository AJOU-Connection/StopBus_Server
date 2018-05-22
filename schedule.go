package main

import (
	"fmt"
	"time"

	"github.com/aubm/interval"
)

func TargetObserver(routeID string, stationID string) {
	isSuccess := false
	stop := interval.Start(func() {
		isSuccess = Observer(routeID, stationID)
	}, 10*time.Second)
	for {
		if isSuccess {
			stop()
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func Observer(routeID string, stationID string) bool {
	fmt.Println("Observer()")
	return true
}
