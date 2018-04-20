package main

import (
	"fmt"
	"log"
	"testing"
)

func TestQuery(t *testing.T) {
	db := Database{}

	ret, err := db.Query("SELECT * FROM Route WHERE number LIKE \"%1%\"")
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		return
	}
	fmt.Println(ret)
}
