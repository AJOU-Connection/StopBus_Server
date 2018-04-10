package main

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	db := Database{"mysql", config.DB.User, config.DB.Password, config.DB.IPAddress, config.DB.Port, config.DB.Name}
	fmt.Println(len(db.Query("SELECT * FROM Route")))
}
