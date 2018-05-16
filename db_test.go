package main

import (
	"log"
	"testing"
)

func TestQuery(t *testing.T) {
	ret := addUserToken("testToken1")
	if ret == -1 {
		log.Printf("expected %v; got %v\n", 1, ret)
	}
}
