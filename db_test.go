package main

import (
	"database/sql"
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
	ret := addGetIn(Reserv{"ffffffff-ecf9-c2fa-4ea2-e8ae34ee3c2b", string([]rune(strconv.Itoa(time.Now().Nanosecond()))[0:9]), string([]rune(strconv.Itoa(time.Now().Nanosecond()))[0:9]), ""})
	if ret != nil {
		log.Printf("expected %v; got %v\n", 1, ret)
	}
}

func TestAddDriverStop(t *testing.T) {
	_, err := addDriverStop(StopInput{"234000026", "234000026"}, GetOff)
	if err != nil {
		log.Printf("expected %v; got %v\n", 1, err)
	}
}

func TestGetGetCount(t *testing.T) {
	_, err := getGetCount("test", "test")
	if err == sql.ErrNoRows {
		// pass
	} else if err != nil {
		log.Printf("%v\n", err)
	}
}
