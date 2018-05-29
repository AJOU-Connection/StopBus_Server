package main

import (
	"fmt"
	"time"

	"github.com/aubm/interval"
)

func isBusPassed(routeID string, stationID string) {
	var prePlateNo, currentPlateNo string

	stop := interval.Start(func() {
		currentPlateNo = getFirstPlateNo(routeID, stationID)
	}, 10*time.Second)

	ticker := time.Tick(5 * time.Second)
	for range ticker {
		if prePlateNo == "" {
			prePlateNo = currentPlateNo
			continue
		}

		if prePlateNo != currentPlateNo {
			deleteDriverStop(routeID, stationID)
			fmt.Println("isBusPassed end")
			stop()
			break
		}
	}
}

func getFirstPlateNo(routeID string, stationID string) string {
	return GetBusArrivalOnlyOne(routeID, stationID).PlateNo1
}

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
