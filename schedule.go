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
			fmt.Println("Alert:", stationID, "/", routeID)
			stop()
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func Observer(routeID string, stationID string) bool {
	arrivalItem := GetBusArrivalOnlyOne(routeID, stationID)
	if arrivalItem.PredictTime1 == 2 {
		return true
	}
	return false
}
