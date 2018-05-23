package main

import "fmt"

// Alert is a function
func GetInAlert(routeID string, stationID string) {
	fmt.Println("곧 버스가 도착합니다.")
	fmt.Println("routeID:", routeID)
	fmt.Println("stationID:", stationID)
}
