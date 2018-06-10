package main

import (
	"fmt"
	"time"

	"github.com/aubm/interval"
)

func isTargetBusPassed(routeID string, stationID string, plateNo string) {
	var currentPlateNo string

	if currentPlateNo = getFirstPlateNo(routeID, stationID); currentPlateNo[len(currentPlateNo)-4:] == plateNo {
		fmt.Println("처음으로 도착하는 버스랑 같음")
		fmt.Println("DB에 Driver를 추가")
		addDriverStop(StopInput{routeID, stationID}, GetOff)
		fmt.Println("버스가 지나가는지 계속 체크 시작")
		isGetOutBusPassed(routeID, stationID, plateNo)
		fmt.Println("종료")
		return
	}
	fmt.Println("처음으로 도착하는 버스랑 다름- 다음 버스를 기다림")
	stop := interval.Start(func() {
		currentPlateNo = getFirstPlateNo(routeID, stationID)
	}, 10*time.Second)

	ticker := time.Tick(5 * time.Second)
	for range ticker {
		if currentPlateNo[len(currentPlateNo)-4:] == plateNo {
			fmt.Println("DB에 Driver를 추가")
			addDriverStop(StopInput{routeID, stationID}, GetOff)
			fmt.Println("버스가 지나가는지 계속 체크 시작")
			isGetOutBusPassed(routeID, stationID, plateNo)
			fmt.Println("종료")
			stop()
		}
	}
}

func isGetOutBusPassed(routeID string, stationID string, plateNo string) {
	var prePlateNo, currentPlateNo string
	isAlert := false

	stop := interval.Start(func() {
		currentPlateNo = getFirstPlateNo(routeID, stationID)
	}, 10*time.Second)

	ticker := time.Tick(5 * time.Second)
	for range ticker {
		if prePlateNo == "" {
			prePlateNo = currentPlateNo
			continue
		}

		arrivalItem := GetBusArrivalOnlyOne(routeID, stationID)
		if (arrivalItem.PredictTime1 <= 2) && (!isAlert) {
			fmt.Println("2분미만으로 사용자에게 알림이 감")
			isAlert = true
			GetOutAlert(routeID, stationID, plateNo)
			fmt.Println("2분미만으로 사용자를 DB에서 지움")
			deleteGetOut(routeID, stationID, plateNo[len(plateNo)-4:])
		}

		if prePlateNo != currentPlateNo {
			fmt.Println("버스가 지나가서 DB에서 정차여부를 삭제함")
			deleteDriverStop(routeID, stationID)
			stop()
			break
		}
		fmt.Println("체크중")
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
	return arrivalItem.PredictTime1 == 2
}
