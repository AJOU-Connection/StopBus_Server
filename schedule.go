package main

import (
	"time"

	"github.com/aubm/interval"
)

func TargetObserver(routeID string, stationID string) {
	isSuccess := false
	stop := interval.Start(func() {
		isSuccess = Observer(routeID, stationID)
	}, 10*time.Second)

	ticker := time.Tick(5 * time.Second)
	for range ticker {
		if isSuccess {
			GetInAlert(routeID, stationID)
			deleteGetIn(routeID, stationID)
			stop()
			break
		}
	}
}

func Observer(routeID string, stationID string) bool {
	arrivalItem := GetBusArrivalOnlyOne(routeID, stationID)
	return arrivalItem.PredictTime1 == 2
}
