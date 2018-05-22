package main

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestAddUserToken(t *testing.T) {
	ret := addUserToken(User{"testToken_" + strconv.Itoa(time.Now().Nanosecond()), "testUUID_" + strconv.Itoa(time.Now().Nanosecond())})
	if ret != nil {
		log.Printf("expected %v; got %v\n", 1, ret)
	}
}

func TestAddGetIn(t *testing.T) {
	ret := addGetIn(GetIn{"testToken_1", "tRouteID1", "StationID"})
	if ret != nil {
		log.Printf("expected %v; got %v\n", 1, ret)
	}
}
