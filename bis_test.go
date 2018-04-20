package main

import (
	"fmt"
	"testing"
)

func TestBusStopNumberToID(t *testing.T) {
	busMobileNumber := "04238"
	ret := BusStopNumberToID(busMobileNumber, 2)
	fmt.Println("bus station mobile number: " + busMobileNumber)
	fmt.Println("bus station ID: " + ret)
}
