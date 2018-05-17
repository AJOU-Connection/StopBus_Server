package main

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestAddUserToken(t *testing.T) {
	ret := addUserToken(User{"testToken_" + strconv.Itoa(time.Now().Nanosecond()), "testUUID_" + strconv.Itoa(time.Now().Nanosecond())})
	if ret == -1 {
		log.Printf("expected %v; got %v\n", 1, ret)
	}
}
