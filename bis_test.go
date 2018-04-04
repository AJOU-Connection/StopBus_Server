package main

import (
	"fmt"
	"testing"
)

func TestBusStopNumberToID(t *testing.T) {
	ret := BusStopNumberToID("04238", 2)
	fmt.Println(ret)
}
