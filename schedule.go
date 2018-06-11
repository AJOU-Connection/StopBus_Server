package main

import (
	"time"

	"github.com/aubm/interval"
)

func isTargetBusPassed(routeID string, stationID string, plateNo string) {
	var currentPlateNo string

	if currentPlateNo = getFirstPlateNo(routeID, stationID); currentPlateNo[len(currentPlateNo)-4:] == plateNo {
		addDriverStop(StopInput{routeID, stationID}, GetOff)
		isGetOutBusPassed(routeID, stationID, plateNo)
		return
	}
	stop := interval.Start(func() {
		currentPlateNo = getFirstPlateNo(routeID, stationID)
	}, 6*time.Second)

	ticker := time.Tick(2 * time.Second)
	for range ticker {
		if currentPlateNo[len(currentPlateNo)-4:] == plateNo {
			addDriverStop(StopInput{routeID, stationID}, GetOff)
			isGetOutBusPassed(routeID, stationID, plateNo)
			stop()
		}
	}
}

func isGetOutBusPassed(routeID string, stationID string, plateNo string) {
	var prePlateNo, currentPlateNo string
	isAlert := false

	stop := interval.Start(func() {
		currentPlateNo = getFirstPlateNo(routeID, stationID)
	}, 6*time.Second)

	ticker := time.Tick(2 * time.Second)
	for range ticker {
		if prePlateNo == "" {
			prePlateNo = currentPlateNo
			continue
		}

		arrivalItem := GetBusArrivalOnlyOne(routeID, stationID)
		if (arrivalItem.PredictTime1 <= 2) && (!isAlert) && (currentPlateNo[len(currentPlateNo)-4:] == plateNo) {
			isAlert = true
			GetOutAlert(routeID, stationID, plateNo)
			deleteGetOut(routeID, stationID, plateNo[len(plateNo)-4:])
		}

		if prePlateNo != currentPlateNo {
			deleteDriverStop(routeID, stationID)
			stop()
			break
		}
	}
}

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
	return arrivalItem.PredictTime1 <= 2
}
